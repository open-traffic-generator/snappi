"""Snappi Generator

Generates a python package based on the
Open Traffic Generator openapi.yaml file.

Generation rules for this file are in GENERATORRULES.md

TBD: 
- response class generation - DONE
- Api return response instance - DONE
- constants - DONE
- parent, choice in child choice classes
- packet slicing using constants
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
        self._generated_methods = []
        self._generated_classes = []
        self._generated_top_level_factories = []
        self._dependencies = ['requests', 'pyyaml', 'jsonpath-ng']
        self._openapi_filename = None
        if 'GITHUB_ACTION' not in os.environ and dependencies is False:
            self._dependencies = []
            self._openapi_filename = openapi_filename
        self.__python = os.path.normpath(sys.executable)
        self.__python_dir = os.path.dirname(self.__python)
        self._src_dir = os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', 'snappi')
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
                'https://github.com/open-traffic-generator/models/releases'
                '/latest/download/openapi.yaml')
            response = requests.request('GET',
                                        OPENAPI_URL,
                                        allow_redirects=True)
            if response.status_code != 200:
                raise Exception(
                    'Unable to retrieve the Open Traffic Generator openapi.yaml'
                    ' file [%s]' % response.content)
            openapi_content = response.content
        else:
            with open(self._openapi_filename, 'rb') as fp:
                openapi_content = fp.read()
        self._openapi = yaml.safe_load(openapi_content)
        self._openapi_version = self._openapi['info']['version']
        print('generating using model version %s' % self._openapi_version)

    def _generate(self):
        self._api_filename = os.path.join(self._src_dir, 'snappi.py')
        # cleanup existing file
        with open(self._api_filename, 'w') as self._fid:
            self._write(0, 'import importlib')

            self._write(0, 'try:')
            self._write(1, 'from typing import Union')
            self._write(0, 'except ImportError:')
            self._write(1, 'pass')

            self._write(0, 'try:')
            self._write(1, 'from typing import Literal')
            self._write(0, 'except ImportError:')
            self._write(1, 'pass')

            self._write(0, 'from .snappicommon import SnappiObject')
            self._write(0, 'from .snappicommon import SnappiList')
            self._write(0, 'from .snappicommon import SnappiHttpTransport')

        methods, factories = self._get_methods_and_factories()

        self._write_api_class(methods, factories)
        self._write_http_api_class(methods)
        self._write_http_server(methods)
        self._write_api_factory()
        self._write_init()
        return self

    def _write_init(self):
        filename = os.path.join(self._src_dir, '__init__.py')
        with open(filename, 'w') as self._fid:
            for class_name in self._generated_classes:
                self._write(0, 'from .snappi import %s' % class_name)
            for factory_name in self._generated_top_level_factories:
                self._write(0, 'from .snappi import %s' % factory_name)

    def _find(self, path, schema_object):
        finds = parse(path).find(schema_object)
        for find in finds:
            yield find.value
            parse(path).find(find.value)

    def _get_methods_and_factories(self):
        """
        Parse methods and top level objects from yaml file to be later used in
        code generation.
        """
        methods = []
        factories = []
        refs = []
        self._top_level_schema_refs = []

        # parse methods
        for path in self._get_api_paths():
            operation = path['operation']
            method_name = operation['operationId'].replace('.', '_').lower()
            if method_name in self._generated_methods:
                continue
            self._generated_methods.append(method_name)
            print('found method %s' % method_name)

            request = parse('$..requestBody..schema').find(operation)
            for req in request:
                _, _, _, ref = self._get_object_property_class_names(req.value)
                if ref:
                    refs.append(ref)

            response = parse('$..responses..schema').find(operation)
            response_object = ''
            response_type = None
            if len(response) == 0:
                # since some responses currently directly $ref to a schema
                # stored someplace else, we need to go one level deeper to
                # get actual response type (currently extracting only for 200)
                response = parse('$..responses.."200"').find(operation)
                response_name, _, _, _ = self._get_object_property_class_names(
                    response[0].value
                )
                response = parse('$.."$ref"').find(
                    self._openapi['components']['responses'][response_name]
                )
                if len(response) > 0:
                    _, response_type, _, ref = self._get_object_property_class_names(
                        response[0].value
                    )
                    if ref:
                        refs.append(ref)
            else:
                _, response_type, _, ref = self._get_object_property_class_names(
                    response[0].value
                )
                if ref:
                    refs.append(ref)

            if response_type is None:
                # TODO: response type is usually None for schema which does not
                # contain any ref (e.g. response of POST /results/capture)
                pass

            methods.append({
                'name': method_name,
                'args': ['self'] if len(request) == 0 else ['self', 'payload'],
                'http_method': path['method'],
                'url': path['url'],
                'description': self._get_description(operation),
                'response_type': response_type
            })

        # parse top level objects (arguments for API requests)
        for ref in refs:
            if ref in self._generated_methods:
                continue
            self._generated_methods.append(ref)
            ret = self._get_object_property_class_names(ref)
            object_name, property_name, class_name, _ = ret
            schema_object = self._get_object_from_ref(ref)
            if 'type' not in schema_object:
                continue
            print('found top level factory method %s' % property_name)
            if schema_object['type'] == 'array':
                ref = schema_object['items']['$ref']
                _, _, class_name, _ = self._get_object_property_class_names(ref)
                class_name = '%sList' % class_name
                self._top_level_schema_refs.append((ref, property_name))
            self._top_level_schema_refs.append((ref, None))

            factories.append({
                'name': property_name,
                'class_name': class_name
            })

        for ref, property_name in self._top_level_schema_refs:
            if property_name is None:
                self._write_snappi_object(ref)
            else:
                self._write_snappi_list(ref, property_name)

        return methods, factories

    def _write_http_api_class(self, methods):
        self._generated_classes.append('HttpApi')
        with open(self._api_filename, 'a') as self._fid:
            self._write()
            self._write()
            self._write(0, 'class HttpApi(Api):')
            self._write(1, '"""%s' % 'Snappi HTTP API')
            self._write(1, '"""')
            self._write(1, 'def __init__(self, host=None):')
            self._write(2, 'super(HttpApi, self).__init__(host=host)')
            self._write(
                2, 'self._transport = SnappiHttpTransport(host=self.host)'
            )

            for method in methods:
                print('generating method %s' % method['name'])
                self._write()
                self._write(
                    1,
                    'def %s(%s):' % (method['name'], ', '.join(method['args']))
                )
                self._write(
                    2,
                    '"""%s %s' % (method['http_method'].upper(), method['url'])
                )
                self._write(0)
                self._write(2, '%s' % method['description'])
                self._write(0)
                self._write(2, 'Return: %s' % method['response_type'])
                self._write(2, '"""')

                self._write(2, 'return self._transport.send_recv(')
                self._write(3, '"%s",' % method['http_method'])
                self._write(3, '"%s",' % method['url'])
                self._write(
                    3, 'payload=%s,' % (
                        method['args'][1] if len(method['args']) > 1 else 'None'
                    )
                )
                self._write(
                    3, 'return_object=%s,' % (
                        'self.' + method['response_type'] + '()' if method[
                            'response_type'
                        ] else 'None'
                    )
                )
                self._write(2, ')')

    def _write_http_server(self, methods):
        pass

    def _write_api_class(self, methods, factories):
        self._generated_classes.append('Api')
        with open(self._api_filename, 'a') as self._fid:
            self._write()
            self._write()
            self._write(0, 'class Api(object):')
            self._write(1, '"""%s' % 'Snappi Abstract API')
            self._write(1, '"""')
            self._write()
            self._write(1, 'def __init__(self, host=None):')
            self._write(2, 'self.host = host if host else "https://localhost"')

            for method in methods:
                print('generating method %s' % method['name'])
                self._write()
                self._write(
                    1,
                    'def %s(%s):' % (method['name'], ', '.join(method['args']))
                )
                self._write(
                    2,
                    '"""%s %s' % (method['http_method'].upper(), method['url'])
                )
                self._write(0)
                self._write(2, '%s' % method['description'])
                self._write(0)
                self._write(2, 'Return: %s' % method['response_type'])
                self._write(2, '"""')
                self._write(
                    2, 'raise NotImplementedError("%s")' % method['name']
                )

            for factory in factories:
                print(
                    'generating top level factory method %s' % factory['name']
                )
                self._write()
                self._write(1, 'def %s(self):' % factory['name'])
                self._write(
                    2, '"""Factory method that creates an instance of %s' % (
                        factory['class_name']
                    )
                )
                self._write()
                self._write(2, 'Return: %s' % factory['class_name'])
                self._write(2, '"""')
                self._write(2, "return %s()" % factory['class_name'])

    def _write_api_factory(self):
        self._generated_top_level_factories.append('api')
        with open(self._api_filename, 'a') as self._fid:
            self._write()
            self._write()
            self._write(0, 'def api(host=None, ext=None):')
            self._write(1, 'if ext is None:')
            self._write(2, 'return HttpApi(host=host)')
            self._write(0)
            self._write(1, 'try:')
            self._write(
                2, 'lib = importlib.import_module("snappi_%s" % ext)'
            )
            self._write(
                2, 'return lib.Api(host=host)'
            )
            self._write(1, 'except ImportError as err:')
            self._write(
                2,
                'msg = "Snappi extension %s is not installed or invalid: %s"'
            )
            self._write(2, 'raise Exception(msg % (ext, err))')

    def _get_object_property_class_names(self, ref):
        object_name = None
        property_name = None
        class_name = None
        ref_name = None
        if isinstance(ref, dict) is True and '$ref' in ref:
            ref_name = ref['$ref']
        elif isinstance(ref, str) is True:
            ref_name = ref
        if ref_name is not None:
            object_name = ref_name.split('/')[-1]
            property_name = object_name.lower().replace('.', '_')
            class_name = object_name.replace('.', '')
        return (object_name, property_name, class_name, ref_name)

    def _write_snappi_object(self, ref, choice_method_name=None):
        schema_object = self._get_object_from_ref(ref)
        ref_name = ref.split('/')[-1]
        class_name = ref_name.replace('.', '')
        if class_name in self._generated_classes:
            return
        self._generated_classes.append(class_name)

        print('generating class %s' % (class_name))
        refs = []
        with open(self._api_filename, 'a') as self._fid:
            self._write()
            self._write()
            self._write(0, 'class %s(SnappiObject):' % class_name)
            slots = ''
            # if choice_method_name is not None:
            slots = "'_parent', '_choice'"
            self._write(1, "__slots__ = (%s)" % slots)
            self._write()

            # write _TYPES definition
            # TODO: this func won't detect whether $ref for a given property is
            # a list because it relies on 'type' attribute to do so
            snappi_types = self._get_snappi_types(schema_object)
            if len(snappi_types) > 0:
                self._write(1, '_TYPES = {')
                for name, value in snappi_types:
                    self._write(2, "'%s': '%s'," % (name, value))
                self._write(1, '}')
                self._write()
            else:
                # TODO: provide empty types as workaround because deserializer
                # in snappicommon.py currently expects it
                self._write(1, '_TYPES = {}')
                self._write()
            
            # write constants
            # search for all simple properties with enum or 
            # x-constant and add them here
            for enum in parse('$..enum | x-constants').find(schema_object):
                for name in enum.value:
                    value = name
                    if isinstance(enum.value, dict):
                        value = enum.value[name]
                    self._write(1, '%s = \'%s\'' % (name.upper(), value))
                if len(enum.value) > 0:
                    self._write()

            # write def __init__(self)
            init_param_string = ''
            # if choice_method_name is not None:
            init_param_string = ", parent=None, choice=None"  # everything will have a parent choice
            for init_param in self._get_simple_type_names(schema_object):
                init_param_string += ', %s=None' % (init_param)
            self._write(1, 'def __init__(self%s):' % init_param_string)
            self._write(2, 'super(%s, self).__init__()' % class_name)
            # if choice_method_name is not None:
            self._write(2, 'self._parent = parent')
            self._write(2, 'self._choice = choice')
            for init_param in self._get_simple_type_names(schema_object):
                self._write(2, "self._set_property('%s', %s)" % (init_param, init_param))
            # if len(parse('$..choice').find(schema_object)) > 0:
            #     self._write(2, 'self.choice = None')

            # process properties - TBD use this one level up to process 
            # schema, in requestBody, Response and also 
            refs = self._process_properties(class_name, schema_object, choice_child=choice_method_name is not None)

        # descend into child properties
        for ref in refs:
            self._write_snappi_object(ref[0], ref[3])
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

    def _process_properties(self, class_name=None, schema_object=None, choice_child=False):
        """Process all properties of a /component/schema object
        Write a factory method for all choice
        If there are no properties then the schema_object is a primitive or array type
        """
        refs = []
        if 'properties' in schema_object:
            choice_names = self._get_choice_names(schema_object)
            excluded_property_names = []
            for choice_name in choice_names:
                if '$ref' not in schema_object['properties'][choice_name]:
                    continue
                ref = schema_object['properties'][choice_name]['$ref']
                self._write_factory_method(None, choice_name, ref)
                excluded_property_names.append(choice_name)
            for property_name in schema_object['properties']:
                if property_name in excluded_property_names:
                    continue
                property = schema_object['properties'][property_name]
                write_set_choice = property_name in choice_names and property_name != 'choice'
                self._write_snappi_property(schema_object, property_name, property, write_set_choice)
            for property_name, property in schema_object['properties'].items():
                ref = parse("$..'$ref'").find(property)
                if len(ref) > 0:
                    restriction = self._get_type_restriction(property)
                    choice_name = property_name if property_name in excluded_property_names else None
                    refs.append((ref[0].value, restriction.startswith('list['), property_name, choice_name))
        return refs

    def _write_snappi_list(self, ref, property_name):
        """This is class writer for schema object properties that have the
        following definition:
        ```
        properties:
          ports:
            type: array
            items:
              $ref: '#/components/schema/...'
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
        contained_class_name = ref_name.replace('.', '')
        class_name = '%sList' % contained_class_name
        if class_name in self._generated_classes:
            return
        self._generated_classes.append(class_name)

        self._imports = []
        print('generating class %s' % (class_name))
        with open(self._api_filename, 'a') as self._fid:
            self._write()
            self._write()
            self._write(0, 'class %s(SnappiList):' % class_name)
            self._write(1, "__slots__ = ()")
            self._write()
            self._write(1, 'def __init__(self):')
            self._write(2, 'super(%s, self).__init__()' % class_name)
            
            # TBD: pass in information to properly construct the __getitem__
            # hint. type: () -> Union[FlowHeader, FlowEthernet, FlowVlan...]
            # might need to be moved the end of this method
            self._write_snappilist_special_methods(contained_class_name, yobject)

            # write factory method for the schema object in the list
            self._write_factory_method(contained_class_name, ref_name.lower().split('.')[-1], ref, True, False)
            # write choice factory methods if the only properties are choice properties
            write_factory_choice_methods = True
            if 'properties' in yobject and 'choice' in yobject['properties']:
                for property in yobject['properties']:
                    if property not in yobject['properties']['choice']['enum']:
                        write_factory_choice_methods = False
                        break
            else:
                write_factory_choice_methods = False                
            if write_factory_choice_methods is True:
                for property in yobject['properties']:
                    if property not in yobject['properties']['choice']['enum']:
                        continue
                    if '$ref' not in yobject['properties'][property]:
                        continue
                    ref = yobject['properties'][property]['$ref']
                    self._write_factory_method(contained_class_name, property, ref, True, True)
        return class_name

    def _write_snappilist_special_methods(self, contained_class_name, schema_object):
        get_item_class_names = [contained_class_name]
        if 'properties' in schema_object and 'choice' in schema_object['properties']:
            for property in schema_object['properties']:
                if property in schema_object['properties']['choice']['enum']:
                    if '$ref' in schema_object['properties'][property]:
                        ref = schema_object['properties'][property]['$ref']
                        _, _, choice_class_name, _ = self._get_object_property_class_names(ref)
                        get_item_class_names.append(choice_class_name)
        get_item_class_names = set(get_item_class_names)
        self._write()
        self._write(1, 'def __getitem__(self, key):')
        self._write(2, '# type: () -> Union[%s]' % (', '.join(get_item_class_names)))
        self._write(2, 'return self._getitem(key)')
        self._write()
        self._write(1, 'def __iter__(self):')
        self._write(2, '# type: () -> %sList' % contained_class_name)
        self._write(2, 'return self._iter()')
        self._write()
        self._write(1, 'def __next__(self):')
        self._write(2, '# type: () -> %s' % contained_class_name)
        self._write(2, 'return self._next()')
        self._write()
        self._write(1, 'def next(self):')
        self._write(2, '# type: () -> %s' % contained_class_name)
        self._write(2, 'return self._next()')

    def _write_factory_method(self,
                              contained_class_name,
                              method_name,
                              ref,
                              snappi_list=False, 
                              choice_method=False):
        yobject = self._get_object_from_ref(ref)
        object_name, property_name, class_name, _ = self._get_object_property_class_names(ref)
        param_string, properties = self._get_property_param_string(yobject)
        self._write()
        if snappi_list is True:
            self._imports.append('from .%s import %s' % (class_name.lower(), class_name))
            self._write(1, 'def %s(self%s):' % (method_name, param_string))
            if contained_class_name is not None:
                self._write(2, "# type: () -> %sList" % (contained_class_name))
            else:
                self._write(2, "# type: () -> %s" % (class_name))
            self._write(2, '"""Factory method that creates an instance of %s class' % (class_name))
            self._write()
            self._write(2, '%s' % self._get_description(yobject))
            self._write(2, '"""')
            if choice_method is True:
                self._write(2, 'item = %s()' % (contained_class_name))
                self._write(2, "item.%s" % (method_name))
                self._write(2, "item.choice = '%s'" % (method_name))
            else:
                params = []
                for property in properties:
                    params.append('%s=%s' % (property, property))
                self._write(2, 'item = %s(%s)' % (class_name, ', '.join(params)))
            self._write(2, 'self._add(item)')
            self._write(2, 'return self')
        else:
            self._write(1, '@property')
            self._write(1, 'def %s(self):' % (method_name))
            self._write(2, "# type: () -> %s" % (class_name))
            self._write(2, '"""Factory property that returns an instance of the %s class' % (class_name))
            self._write()
            self._write(2, '%s' % self._get_description(yobject))
            self._write(2, '"""')
            self._write(2, "return self._get_property('%s', %s(self, '%s'))" % (method_name, class_name, method_name))

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

    def _write_snappi_property(self, schema_object, name, property, write_set_choice=False):
        ref = parse("$..'$ref'").find(property)
        restriction = self._get_type_restriction(property)
        if len(ref) > 0:
            object_name = ref[0].value.split('/')[-1]
            class_name = object_name.replace('.', '')
            if restriction.startswith('list['):
                type_name = '%sList' % class_name
            else:
                type_name = class_name
        else:
            type_name = restriction
        self._write()
        self._write(1, '@property')
        self._write(1, 'def %s(self):' % name)
        self._write(2, "# type: () -> %s" % (type_name))
        self._write(2, '"""%s getter' % (name))
        self._write()
        self._write(2, self._get_description(property))
        self._write()
        self._write(2, 'Returns: %s' % restriction)
        self._write(2, '"""')
        if len(parse("$..'type'").find(property)) > 0 and len(ref) == 0:
            self._write(2, "return self._get_property('%s')" % (name))
            self._write()
            self._write(1, '@%s.setter' % name)
            self._write(1, 'def %s(self, value):' % name)
            self._write(2, '"""%s setter' % (name))
            self._write()
            self._write(2, self._get_description(property))
            self._write()
            self._write(2, 'value: %s' % restriction)
            self._write(2, '"""')
            if write_set_choice is True:
                self._write(2, "self._set_property('%s', value, '%s')" % (name, name))
            else:
                self._write(2, "self._set_property('%s', value)" % (name))
        elif len(ref) > 0:
            if restriction.startswith('list['):
                self._write(2, "return self._get_property('%s', %sList)" % (name, class_name))
            else:
                self._write(2, "return self._get_property('%s', %s)" % (name, class_name))

    def _get_description(self, yobject):
        if 'description' not in yobject:
            yobject['description'] = 'TBD'
        # remove tabs, multiple spaces
        description = re.sub('\n', '. ', yobject['description'])
        description = re.sub('\s+', ' ', description)
        return description
        # doc_string = []
        # for line in re.split('\. ', description):
        #     line = re.sub('\.$', '', line)
        #     if len(line) > 0:
        #         doc_string.append('%s  ' % line)
        # return doc_string

    def _get_snappi_types(self, yobject):
        types = []
        if 'properties' in yobject:
            for name in yobject['properties']:
                yproperty = yobject['properties'][name]
                ref = parse("$..'$ref'").find(yproperty)
                if len(ref) > 0:
                    object_name = ref[0].value.split('/')[-1]
                    class_name = object_name.replace('.', '')
                    if 'type' in yproperty and yproperty['type'] == 'array':
                        class_name += 'List'
                    types.append((name, class_name))
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
                        'operation': yobject[method],
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
        leaf = self._openapi
        for attr in ref.split('/')[1:]:
            leaf = leaf[attr]
        return leaf

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
