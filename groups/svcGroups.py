import os
import sys
import grpc
from flask import g

import opentracing
from grpc_opentracing import open_tracing_client_interceptor
from grpc_opentracing.grpcext import intercept_channel


path = os.path.dirname(os.path.abspath(__file__))
sys.path.append(path)

from groups_pb2_grpc import groupsStub
from groups_pb2 import Group

class SvcGroups():
  def __init__(self, url):
    print("connect to groups svc")
    channel = grpc.insecure_channel(url) # os.environ['GROUPS_URL']
    interceptor = open_tracing_client_interceptor(opentracing.tracer)
    channel = intercept_channel(channel, interceptor)
    self.stub = groupsStub(channel)
    print("connected :-)")

  def GetPaths(self, bearer, identity):
    metadata = (('authorization', bearer),
            ('some-md-key', 'another value'))
    grp = self.stub.Get(request=Group(name=identity), metadata=metadata)
    print(grp.id)
    if grp.id == 0 and grp.name != "root":
      return [], False
    return grp.paths, True
