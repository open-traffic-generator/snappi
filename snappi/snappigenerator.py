"""Snappi Generator

Generates a python package based on the
Open Traffic Generator openapi.yaml file.

Generation rules for this file are in GENERATORRULES.md

TBD: 
- response class generation
- Api return response instance
- constants
- slicing
- docstrings
- type checking
"""
import sys
import yaml
import os
import subprocess
import re
import requests
from jsonpath_ng import parse


class SnappiGenerator(object):
    """Builds the snappi python package based on a released version of the
    open-traffic-generator openapi.yaml file.
    """
    def __init__(self, dependencies=True, openapi_filename=None):
        self._dependencies = ['requests', 'pyyaml', 'jsonpath-ng']
        if 'GITHUB_ACTION' not in os.environ and dependencies is False:
            self._dependencies = []
        self._openapi_filename = openapi_filename
        self.__python = os.path.normpath(sys.executable)
        self.__python_dir = os.path.dirname(self.__python)
        self._src_dir = os.path.dirname(os.path.abspath(__file__))
        self._docs_dir = os.path.join(self._src_dir, '..', 'docs')
        self._clean()
        self._install_dependencies()
        self._get_openapi_file()
        self._generate()

    def _clean(self):
        """Clean the environment prior to file generation
        - Remove any locally installed version of snappi
        - Remove generated files
        - Leave infrastructure files that are prefixed with the word snappi
        """
        process_args = [
            self.__python, '-m', 'pip', 'uninstall', '--yes', 'snappi'
        ]
        subprocess.Popen(process_args, shell=False).wait()
        import os
        import fnmatch
        for rootDir, subdirs, filenames in os.walk(self._src_dir):
            if rootDir.endswith('tests'):
                continue
            for filename in fnmatch.filter(filenames, '*.py'):
                try:
                    if filename.startswith('snappi') is False:
                        os.remove(os.path.join(rootDir, filename))
                except OSError:
                    print('Error deleting file %s' % filename)

    def _install_dependencies(self):
        """Install any dependencies this project requires
        """
        for package in self._dependencies:
            print('installing dependency %s...' % package)
            process_args = [
                self.__python, '-m', 'pip', 'install', '--upgrade', package
            ]
            subprocess.Popen(process_args, shell=False).wait()

    def _get_openapi_file(self):
        if self._openapi_filename is None:
            OPENAPI_URL = (
                'https://github.com/open-traffic-generator/models/releases/latest/download'
                '/openapi.yaml')
            response = requests.request('GET',
                                        OPENAPI_URL,
                                        allow_redirects=True)
            if response.status_code != 200:
                raise Exception(
                    'Unable to retrieve the Open Traffic Generator openapi.yaml file [%s]'
                    % response.content)
            openapi_content = response.content
        else:
            with open(self._openapi_filename, 'rb') as fp:
                openapi_content = fp.read()
        self._openapi = yaml.safe_load(openapi_content)

    def _generate(self):
        self._write_paths()
        self._write_init()
        return self

    def _write_init(self):
        filename = os.path.join(self._src_dir, '__init__.py')
        with open(filename, 'w') as self._fid:
            self._write(0, 'from .api import Api')

    def _write_paths(self):
        api_filename = os.path.join(self._src_dir, 'api.py')
        with open(api_filename, 'a') as self._fid:
            self._write(0, 'from .snappicommon import SnappiRestTransport')
            self._write()
            self._write()
            self._write(0, 'class Api(SnappiRestTransport):')
            self._write(1, '"""%s' % 'Snappi Abstract API')
            self._write(1, '"""')
            self._write(1, 'def __init__(self):')
            self._write(2, 'super(Api, self).__init__()')

        # write methods
        for path in self._get_api_paths():
            operation = path['operation']
            method_name = operation['operationId'].replace('.', '_').lower()
            print('generating %s in file %s...' % (method_name, api_filename))
            with open(api_filename, 'a') as self._fid:
                self._write()
                self._write(1, 'def %s(self, content=None):' % method_name)
                self._write(2, '"""%s' % 'TBD')
                self._write(2, '"""')
                self._write(
                    2, "return self.send_recv('%s', '%s', payload=content)" %
                    (path['method'], path['url']))

        # write top level objects for requests
        for yobject in self._openapi['paths'].values():
            with open(api_filename, 'a') as self._fid:
                find = []
                for section in ['requestBody', 'responses']:
                    find += parse('$..%s..schema' % section).find(yobject)
                if len(find) == 0:
                    continue
                top_level_schema = find[0].value
                object_name = top_level_schema['$ref'].split('/')[-1]
                property_name = object_name.lower().replace('.', '_')
                class_name = object_name.replace('.', '')
                self._write()
                self._write(1, 'def %s(self):' % property_name)
                self._write(
                    2, '"""%s' %
                    'Return instance of auto-generated top level class %s' %
                    class_name)
                self._write(2, '"""')
                self._write(
                    2, "from .%s import %s" % (class_name.lower(), class_name))
                self._write(2, "return %s()" % (class_name))
            self._write_snappi_object(top_level_schema['$ref'])

    def _write_snappi_object(self, ref):
        schema_object = self._get_object_from_ref(ref)
        ref_name = ref.split('/')[-1]
        class_name = ref_name.replace('.', '')
        class_filename = os.path.join(self._src_dir,
                                      '%s.py' % class_name.lower())
        refs = []
        if os.path.exists(class_filename) is False:
            print('generating %s in file %s...' % (class_name, class_filename))
            with open(class_filename, 'a') as self._fid:
                self._write(0, 'from .snappicommon import SnappiObject')
                self._write()
                self._write()
                self._write(0, 'class %s(SnappiObject):' % class_name)

                # write _TYPES definition
                snappi_types = self._get_snappi_types(schema_object)
                if len(snappi_types) > 0:
                    self._write(1, '_TYPES = {')
                    for name, value in snappi_types:
                        self._write(2, "'%s': '%s'," % (name, value))
                    self._write(1, '}')
                    self._write()

                # write def __init__(self)
                init_param_string = ''
                for init_param in self._get_simple_type_names(schema_object):
                    init_param_string += ', %s=None' % (init_param)
                self._write(1, 'def __init__(self%s):' % init_param_string)
                self._write(2, 'super(%s, self).__init__()' % class_name)
                for init_param in self._get_simple_type_names(schema_object):
                    self._write(2, 'self.%s = %s' % (init_param, init_param))

                # process properties
                refs = self._process_properties(class_name, schema_object)

            # descend into child properties
            for ref in refs:
                self._write_snappi_object(ref[0])
                if ref[1] is True:
                    self._write_snappi_list(ref[0], ref[2])

    def _get_simple_type_names(self, schema_object):
        simple_type_names = []
        if 'properties' in schema_object:
            choice_names = self._get_choice_names(schema_object)
            for name in schema_object['properties']:
                if name in choice_names:
                    continue
                ref = parse("$..'$ref'").find(
                    schema_object['properties'][name])
                if len(ref) == 0:
                    simple_type_names.append(name)
        return simple_type_names

    def _get_choice_names(self, schema_object):
        choice_names = []
        if 'choice' in schema_object['properties']:
            choice_names = schema_object['properties']['choice']['enum']
            choice_names.append('choice')
        return choice_names

    def _process_properties(self, class_name=None, schema_object=None):
        """Process all properties of a /component/schema object
        Write a factory method for all choice
        """
        refs = []
        if 'properties' in schema_object:
            excluded_property_names = []
            for choice_name in self._get_choice_names(schema_object):
                if '$ref' not in schema_object['properties'][choice_name]:
                    continue
                ref = schema_object['properties'][choice_name]['$ref']
                self._write_factory_method(None, choice_name, ref)
                excluded_property_names.append(choice_name)
            for property_name in schema_object['properties']:
                if property_name in excluded_property_names:
                    continue
                property = schema_object['properties'][property_name]
                self._write_snappi_property(schema_object, property_name,
                                            property)
            for property_name, property in schema_object['properties'].items():
                ref = parse("$..'$ref'").find(property)
                if len(ref) > 0:
                    restriction = self._get_type_restriction(property)
                    refs.append((ref[0].value, restriction.startswith('list['),
                                 property_name))
        return refs

    def _write_snappi_list(self, ref, property_name):
        """This is class writer for schema object properties that have the
        following definition:
        ```
        type: array
        items:
            $ref: '#/components/schema/...
        ```

        If the schema object has a property named choice, that property needs 
        to be brought forward so that the generated class can provide factory
        methods for objects for each of the choice $refs (if any).

        if choice exists: 
            for each choice enum:
                generate a factory method named after the choice
                in the method set the choice property
        """
        yobject = self._get_object_from_ref(ref)
        ref_name = ref.split('/')[-1]
        class_name = ref_name.replace('.', '')
        class_filename = os.path.join(self._src_dir,
                                      '%slist.py' % class_name.lower())
        if os.path.exists(class_filename) is False:
            print('generating %s in file %s...' % (class_name, class_filename))
            with open(class_filename, 'a') as self._fid:
                self._write(0, 'from .snappicommon import SnappiList')
                self._write()
                self._write()
                self._write(0, 'class %sList(SnappiList):' % class_name)
                self._write(1, 'def __init__(self):')
                self._write(2, 'super(%sList, self).__init__()' % class_name)
                self._write()
                # write factory method for the schema object in the list
                self._write_factory_method(None, class_name.lower(), ref, True)
                # write choice factory methods if any
                if 'choice' in yobject['properties']:
                    for property in yobject['properties']:
                        if property not in yobject['properties']['choice'][
                                'enum']:
                            continue
                        if '$ref' not in yobject['properties'][property]:
                            continue
                        ref = yobject['properties'][property]['$ref']
                        self._write_factory_method(class_name, property, ref,
                                                   True)

    def _write_factory_method(self,
                              container_class_name,
                              method_name,
                              ref,
                              snappi_list=False):
        yobject = self._get_object_from_ref(ref)
        ref_name = ref.split('/')[-1]
        class_name = ref_name.replace('.', '')
        param_string, properties = self._get_property_param_string(yobject)
        self._write()
        if snappi_list is True:
            self._write(1, 'def %s(self%s):' % (method_name, param_string))
            if container_class_name is not None:
                self._write(
                    2, 'from .%s import %s' %
                    (container_class_name.lower(), container_class_name))
                self._write(2, 'item = %s()' % (container_class_name))
                self._write(2, 'item.%s' % (method_name))
            else:
                self._write(
                    2, 'from .%s import %s' % (class_name.lower(), class_name))
                self._write(
                    2, 'item = %s(%s)' % (class_name, ', '.join(properties)))
            self._write(2, 'self._add(item)')
            self._write(2, 'return self')
        else:
            self._write(1, '@property')
            self._write(1, 'def %s(self):' % (method_name))
            self._write(
                2, "from .%s import %s" % (class_name.lower(), class_name))
            self._write(
                2,
                "if '%s' not in self._properties or self._properties['%s'] is None:"
                % (method_name, method_name))
            self._write(
                3, "self._properties['%s'] = %s()" % (method_name, class_name))
            self._write(2, 'self.choice = \'%s\'' % (method_name))
            # for property_name in properties:
            #     self._write(2, 'self._properties[\'%s\'].%s = %s' % (method_name, property_name, property_name))
            self._write(2, "return self._properties['%s']" % (method_name))

    def _get_property_param_string(self, yobject):
        property_param_string = ''
        properties = []
        if 'properties' in yobject:
            for name, property in yobject['properties'].items():
                if name == 'choice':
                    continue
                default = 'None'
                if 'obj' not in self._get_type_restriction(property):
                    property_param_string += ', %s' % name
                    properties.append(name)
                    if 'default' in property:
                        default = property['default']
                    if 'enum' in property:
                        property_param_string += "='%s'" % default
                    else:
                        property_param_string += '=%s' % default
        return (property_param_string, properties)

    def _write_snappi_property(self, schema_object, name, property):
        ref = parse("$..'$ref'").find(property)
        self._write()
        self._write(1, '@property')
        self._write(1, 'def %s(self):' % name)
        self._write(2, '"""%s getter' % (name))
        self._write()
        for line in self._get_description(property):
            self._write(2, line)
        self._write()
        self._write(2, 'Returns: %s' % self._get_type_restriction(property))
        self._write(2, '"""')
        if len(parse("$..'type'").find(property)) > 0 and len(ref) == 0:
            self._write(2, "return self._properties['%s']" % (name))
            self._write()
            self._write(1, '@%s.setter' % name)
            self._write(1, 'def %s(self, value):' % name)
            self._write(2, '"""%s setter' % (name))
            self._write()
            for line in self._get_description(property):
                self._write(2, line)
            self._write()
            self._write(2, 'value: %s' % self._get_type_restriction(property))
            self._write(2, '"""')
            if name in self._get_choice_names(
                    schema_object) and name != 'choice':
                self._write(2, "self._properties['choice'] = '%s'" % (name))
            self._write(2, "self._properties['%s'] = value" % (name))
        elif len(ref) > 0:
            object_name = ref[0].value.split('/')[-1]
            class_name = object_name.replace('.', '')
            file_name = class_name.lower()
            restriction = self._get_type_restriction(property)
            if restriction.startswith('list['):
                self._write(
                    2, "from .%slist import %sList" %
                    (class_name.lower(), class_name))
                self._write(
                    2,
                    "if '%s' not in self._properties or self._properties['%s'] is None:"
                    % (name, name))
                self._write(
                    3,
                    "self._properties['%s'] = %sList()" % (name, class_name))
                self._write(2, "return self._properties['%s']" % (name))
            else:
                self._write(2, "from .%s import %s" % (file_name, class_name))
                self._write(
                    2,
                    "if '%s' not in self._properties or self._properties['%s'] is None:"
                    % (name, name))
                self._write(
                    3, "self._properties['%s'] = %s()" % (name, class_name))
                self._write(2, "return self._properties['%s']" % (name))

    def _get_description(self, yobject):
        if 'description' not in yobject:
            yobject['description'] = 'TBD'
        # remove tabs, multiple spaces
        description = re.sub('\n', '. ', yobject['description'])
        description = re.sub('\s+', ' ', description)
        doc_string = []
        for line in re.split('\. ', description):
            line = re.sub('\.$', '', line)
            if len(line) > 0:
                doc_string.append('%s  ' % line)
        return doc_string

    def _get_snappi_types(self, yobject):
        types = []
        if 'properties' in yobject:
            for name in yobject['properties']:
                yproperty = yobject['properties'][name]
                ref = parse("$..'$ref'").find(yproperty)
                if len(ref) > 0:
                    object_name = ref[0].value.split('/')[-1]
                    class_name = object_name.replace('.', '')
                    file_name = class_name.lower()
                    if 'type' in yproperty and yproperty['type'] == 'array':
                        class_name += 'List'
                        file_name += 'list'
                    types.append((name, '.%s.%s' % (file_name, class_name)))
        return types

    def _get_default_value(self, property):
        if 'default' in property:
            return property['default']
        property_type = property['type']
        if property_type == 'array':
            return '[]'
        if property_type == 'string':
            return "''"
        if property_type == 'integer':
            return 0
        if property_type == 'number':
            return 0
        if property_type == 'boolean':
            return False
        raise Exception('Missing handler for property type `%s`' %
                        property_type)

    def _get_api_paths(self):
        paths = []
        for url, yobject in self._openapi['paths'].items():
            for method in yobject:
                if method.lower() in ['get', 'post', 'put', 'patch', 'delete']:
                    paths.append({
                        'url': url,
                        'method': method,
                        'operation': yobject[method]
                    })
        return paths

    def _write_component_schemas(self):
        for key, yobject in self._openapi['components']['schemas'].items():
            pieces = key.split('.')
            self._classname = key
            path = self._src_dir + '/'
            if '.' in key:
                self._classname = pieces[-1]
                path += '_'.join(pieces[0:-1]).lower()
            else:
                path += self._classname.lower()
            self._classfilename = path
            print('generating %s in file %s...' %
                  (self._classname, self._classfilename))

            with open(self._classfilename + '.py', 'a',
                      newline='\n') as self._fid:
                self._write()
                self._write()
                self._write(0, 'class %s(object):' % self._classname)

                # create a list of any choice tuples
                choice_tuples = []
                if 'properties' in yobject and 'choice' in yobject[
                        'properties']:
                    if 'required' in yobject and 'choice' not in yobject[
                            'required']:
                        choice_tuples.append(
                            ('None', choice_enum, choice_enum))
                    for choice_enum in yobject['properties']['choice']['enum']:
                        if choice_enum not in yobject['properties']:
                            choice_tuples.append(
                                (choice_enum, choice_enum, None))
                        else:
                            choice = yobject['properties'][choice_enum]
                            if '$ref' in choice:
                                choice_classname = self._get_classname_from_ref(
                                    choice['$ref'])
                                choice_tuples.append(
                                    (choice_classname, choice_enum,
                                     choice['$ref']))
                            elif choice['type'] == 'string':
                                choice_tuples.append(
                                    ('str', choice_enum, None))
                            elif choice['type'] in ['number', 'integer']:
                                choice_tuples.append(
                                    ('float', choice_enum, None))
                                choice_tuples.append(
                                    ('int', choice_enum, None))
                            elif choice['type'] == 'array':
                                choice_tuples.append(
                                    ('list', choice_enum, None))
                            elif choice['type'] == 'boolean':
                                choice_tuples.append(
                                    ('boolean', choice_enum, None))

                # class documentation
                self._write(
                    1,
                    '"""Generated from OpenAPI schema object #/components/schemas/%s'
                    % key)
                self._write()
                if 'description' not in yobject:
                    yobject['description'] = 'TBD'
                # remove tabs, multiple spaces
                description = re.sub('\n', '. ', yobject['description'])
                description = re.sub('\s+', ' ', description)
                for line in re.split('\. ', description):
                    line = re.sub('\.$', '', line)
                    if len(line) > 0:
                        self._write(1, '%s  ' % line)
                if 'properties' in yobject:
                    self._write()
                    self._write(1, "Args")
                    self._write(1, "----")
                    for name, property in yobject['properties'].items():
                        if len([
                                item
                                for item in choice_tuples if item[1] == name
                        ]) > 0:
                            continue
                        if name == 'choice':
                            type = 'Union[%s]' % ', '.join(
                                [item[0] for item in choice_tuples])
                        elif name == 'additionalProperties':
                            name = 'additional_properties'
                            type = '**additional_properties'
                        else:
                            type = self._get_type_restriction(property)
                        if 'description' not in property:
                            property['description'] = 'TBD'
                        description = re.sub('\n', '. ',
                                             property['description'])
                        description = re.sub('\s+', ' ',
                                             property['description'])
                        lines = re.split('\.', description)
                        self._write(
                            1,
                            "- %s (%s): %s" % (name, type, lines[0].strip()))
                        for line in lines[1:]:
                            line = line.strip()
                            if len(line) > 0:
                                self._write(1, ' %s' % line.strip())
                self._write(1, '"""')

                # constants
                if 'x-constants' in yobject.keys():
                    self._write(1)
                    for constant, value in yobject['x-constants'].items():
                        self._write(1, "%s = '%s'" % (constant.upper(), value))
                    self._write(1)

                # choice map
                if len(choice_tuples) > 0:
                    self._write(1, '_CHOICE_MAP = {')
                    for choice_tuple in choice_tuples:
                        self._write(
                            2,
                            "'%s': '%s'," % (choice_tuple[0], choice_tuple[1]))
                    self._write(1, '}')

                # init args
                args = ''
                if 'properties' in yobject:
                    for name, property in yobject['properties'].items():
                        if len([
                                item for item in choice_tuples
                                if item[1] == name
                        ]) == 0:
                            arg_type = 'None'
                            if 'type' in property and property[
                                    'type'] == 'array':
                                arg_type = '[]'
                            elif 'default' in property:
                                if property['type'] == 'string':
                                    arg_type = "'%s'" % property['default']
                                else:
                                    arg_type = '%s' % property['default']
                            args += '%s%s=%s' % (', ', name, arg_type)
                self._write(1, 'def __init__(self%s):' % args)
                if len(args) == 0:
                    self._write(2, 'pass')
                self._write_data_properties(yobject, self._classname,
                                            choice_tuples)
        return self

    def _write_data_properties(self, schema, classname, choice_tuples):
        import_lines = []
        if len(choice_tuples) > 0:
            for choice_tuple in choice_tuples:
                if choice_tuple[2] is not None:
                    import_line = self._get_import_from_ref(choice_tuple[2])
                    if import_line not in import_lines:
                        self._write(2, import_line)
                        import_lines.append(import_line)
            choices = []
            for choice_tuple in choice_tuples:
                choices.append(choice_tuple[0])
            self._write(
                2,
                'if isinstance(choice, (%s)) is False:' % (', '.join(choices)))
            self._write(
                3, "raise TypeError('choice must be of type: %s')" %
                (', '.join(choices)))
            self._write(
                2,
                "self.__setattr__('choice', %s._CHOICE_MAP[type(choice).__name__])"
                % classname)
            self._write(
                2,
                "self.__setattr__(%s._CHOICE_MAP[type(choice).__name__], choice)"
                % classname)

        if 'properties' in schema:
            for name, property in schema['properties'].items():
                if '$ref' in property:
                    import_line = self._get_import_from_ref(property['$ref'])
                    if import_line not in import_lines:
                        self._write(2, import_line)
                        import_lines.append(import_line)
            for name, property in schema['properties'].items():
                if len([item for item in choice_tuples if item[1] == name
                        ]) == 0 and name != 'choice':
                    restriction = self._get_isinstance_restriction(
                        schema, name, property)
                    self._write(
                        2,
                        'if isinstance(%s, %s) is True:' % (name, restriction))
                    if restriction == '(list, type(None))':
                        self._write(
                            3, 'self.%s = [] if %s is None else list(%s)' %
                            (name, name, name))
                    else:
                        if 'pattern' in property:
                            self._write(3, 'import re')
                            self._write(
                                3,
                                "assert(bool(re.match(r'%s', %s)) is True)" %
                                (property['pattern'], name))
                        self._write(3, 'self.%s = %s' % (name, name))
                    self._write(2, 'else:')
                    self._write(
                        3, "raise TypeError('%s must be an instance of %s')" %
                        (name, restriction))

    def _get_isinstance_restriction(self, schema, name, property):
        type_none = ', type(None)'
        if 'required' in schema and name in schema['required']:
            type_none = ''
        if '$ref' in property:
            return '(%s%s)' % (self._get_classname_from_ref(
                property['$ref']), type_none)
        elif name == 'additionalProperties':
            return '**additional_properties'
        elif property['type'] in ['number', 'integer']:
            return '(float, int%s)' % type_none
        elif property['type'] == 'string':
            return '(str%s)' % type_none
        elif property['type'] == 'array':
            return '(list%s)' % type_none
        elif property['type'] == 'boolean':
            return '(bool%s)' % type_none

    def _get_type_restriction(self, property):
        if '$ref' in property:
            ref_obj = self._get_object_from_ref(property['$ref'])
            description = ''
            if 'description' in ref_obj:
                description = ref_obj['description']
            if 'description' in property:
                description += property['description']
            property['description'] = description
            class_name = property['$ref'].split('/')[-1].replace('.', '')
            return 'obj(snappi.%s)' % class_name
        elif 'oneOf' in property:
            return 'Union[%s]' % ','.join(
                [item['type'] for item in property['oneOf']])
        elif property['type'] == 'number':
            return 'float'
        elif property['type'] == 'integer':
            return 'int'
        elif property['type'] == 'string':
            if 'enum' in property:
                return 'Union[%s]' % ', '.join(property['enum'])
            else:
                return 'str'
        elif property['type'] == 'array':
            return 'list[%s]' % self._get_type_restriction(property['items'])
        elif property['type'] == 'boolean':
            return 'boolean'

    def _get_object_from_ref(self, ref):
        pieces = ref.split('/')
        json_path = '$.%s."%s"' % ('.'.join(pieces[1:-1]), pieces[-1])
        return parse(json_path).find(self._openapi)[0].value

    def _get_import_from_ref(self, ref):
        filename = '_'.join(
            ref.lower().split('#/components/schemas/')[-1].split('.')[0:-1])
        classname = self._get_classname_from_ref(ref)
        if len(filename) == 0:
            filename = classname.lower()
        return 'from abstract_open_traffic_generator.%s import %s' % (
            filename, classname)

    def _get_classname_from_ref(self, ref):
        final_piece = ref.split('/')[-1]
        if '.' in final_piece:
            return final_piece.split('.')[-1]
        else:
            return final_piece

    def _write(self, indent=0, line=''):
        self._fid.write('    ' * indent + line + '\n')

    def _bundle(self, base_dir, api_filename, output_filename):
        print('bundling started')
        self._read_file(base_dir, api_filename)
        with open(self._output_filename, 'w') as fid:
            yaml.dump(self._content, fid, indent=2, sort_keys=False)
        print('bundling complete')

    def _read_file(self, base_dir, filename):
        from yaml import safe_load
        filename = os.path.join(base_dir, filename)
        filename = os.path.abspath(os.path.normpath(filename))
        base_dir = os.path.dirname(filename)
        with open(filename) as fid:
            yobject = safe_load(fid)
        self._process_yaml_object(base_dir, yobject)

    def _process_yaml_object(self, base_dir, yobject):
        for key, value in yobject.items():
            if key in ['openapi', 'info', 'servers'
                       ] and key not in self._content.keys():
                self._content[key] = value
            elif key in ['paths']:
                if key not in self._content.keys():
                    self._content[key] = {}
                for sub_key in value.keys():
                    self._content[key][sub_key] = value[sub_key]
            elif key == 'components':
                if key not in self._content.keys():
                    self._content[key] = {'schemas': {}}
                if 'schemas' in value:
                    schemas = value['schemas']
                    for schema_key in schemas.keys():
                        self._content['components']['schemas'][
                            schema_key] = schemas[schema_key]
        self._resolve_refs(base_dir, yobject)

    def _resolve_refs(self, base_dir, yobject):
        """Resolving references is relative to the current file location
        """
        if isinstance(yobject, dict):
            for key, value in yobject.items():
                if key == '$ref' and value.startswith('#') is False:
                    refs = value.split('#')
                    print('resolving %s' % value)
                    self._read_file(base_dir, refs[0])
                    yobject[key] = '#%s' % refs[1]
                elif isinstance(value, str) and 'x-inline' in value:
                    refs = value.split('#')
                    print('inlining %s' % value)
                    inline = self._get_inline_ref(base_dir, refs[0], refs[1])
                    yobject[key] = inline
                else:
                    self._resolve_refs(base_dir, value)
        elif isinstance(yobject, list):
            for item in yobject:
                self._resolve_refs(base_dir, item)

    def _get_inline_ref(self, base_dir, filename, inline_key):
        filename = os.path.join(base_dir, filename)
        filename = os.path.abspath(os.path.normpath(filename))
        base_dir = os.path.dirname(filename)
        with open(filename) as fid:
            yobject = yaml.safe_load(fid)
        return parse('$%s' %
                     inline_key.replace('/', '.'), ).find(yobject)[0].value


if __name__ == '__main__':
    openapi_filename = None
    # openapi_filename = os.path.normpath('../../models/openapi.yaml')
    SnappiGenerator(dependencies=False, openapi_filename=openapi_filename)