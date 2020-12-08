import pytest


@pytest.fixture(scope='session')
def api():
    from ..api import Api
    return Api()
