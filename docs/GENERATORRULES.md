# Generator Rules
The following are the rules that govern how openapi.yaml will 
be translated into snappi objects and snappi lists.

## `paths:`
Given OpenAPI paths...
```yaml
paths:
  /config:
    post:
      operationId: set_config
    get:
      operationId: get_config
```
The generator should do the following:
- Generate a single snappi.api.Api class that encapsulates all operations.
- If the `operationId` is not specified the method and path will become the operationId.
- The / separator character in the path will be replaced with an _ and the method name
will be in lower case.
    ```python
    class Api(object):
        def set_config(self, content):
            pass
        def get_config(self):
            pass
    ```

## `type: object`
Given a /components/schemas object that has a type of object...
```yaml
components:
  schemas:
    Config:
      type: object
```
The generator should do the following:
- Generate a class that inherits from SnappiObject.
- Class name is camel case with any ._- characters removed.
- Class name must be unique.
- Class must be placed in its own file named as lower case class name file.
    ```python
    class Config(SnappiObject):
    ```

## `array of $ref`
Given a property that has a type of array of $ref...
```yaml
components:
  schemas:
    Config:
      type: object
      properties:
        flows:
          type: array
          items:
            $ref: '/components/schemas/Flow'
```
The generator should do the following:
- Generate a property on the container class that returns a SnappiList object.
    ```python
    class Config(SnappiObject):
        @property
        def flows(self):
            from .flowlist import FlowList
            if 'flows' not in self._properties or self._properties['flows'] is None:
                self._properties['flows'] = FlowList()
            return self._properties['flows']
    ```
- Generate a container class that inherits SnappList class.
    ```python
    class FlowList(SnappiList):
    ```


## composite type
```
type: array
items:
  $ref: '/components/schemas/Flow.Header'
```




