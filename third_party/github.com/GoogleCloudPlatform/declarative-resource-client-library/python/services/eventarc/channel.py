# Copyright 2024 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.eventarc import channel_pb2
from google3.cloud.graphite.mmv2.services.google.eventarc import channel_pb2_grpc

from typing import List


class Channel(object):
    def __init__(
        self,
        name: str = None,
        uid: str = None,
        create_time: str = None,
        update_time: str = None,
        third_party_provider: str = None,
        pubsub_topic: str = None,
        state: str = None,
        activation_token: str = None,
        crypto_key_name: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.third_party_provider = third_party_provider
        self.crypto_key_name = crypto_key_name
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = channel_pb2_grpc.EventarcChannelServiceStub(channel.Channel())
        request = channel_pb2.ApplyEventarcChannelRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.third_party_provider):
            request.resource.third_party_provider = Primitive.to_proto(
                self.third_party_provider
            )

        if Primitive.to_proto(self.crypto_key_name):
            request.resource.crypto_key_name = Primitive.to_proto(self.crypto_key_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyEventarcChannel(request)
        self.name = Primitive.from_proto(response.name)
        self.uid = Primitive.from_proto(response.uid)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.third_party_provider = Primitive.from_proto(response.third_party_provider)
        self.pubsub_topic = Primitive.from_proto(response.pubsub_topic)
        self.state = ChannelStateEnum.from_proto(response.state)
        self.activation_token = Primitive.from_proto(response.activation_token)
        self.crypto_key_name = Primitive.from_proto(response.crypto_key_name)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = channel_pb2_grpc.EventarcChannelServiceStub(channel.Channel())
        request = channel_pb2.DeleteEventarcChannelRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.third_party_provider):
            request.resource.third_party_provider = Primitive.to_proto(
                self.third_party_provider
            )

        if Primitive.to_proto(self.crypto_key_name):
            request.resource.crypto_key_name = Primitive.to_proto(self.crypto_key_name)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteEventarcChannel(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = channel_pb2_grpc.EventarcChannelServiceStub(channel.Channel())
        request = channel_pb2.ListEventarcChannelRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListEventarcChannel(request).items

    def to_proto(self):
        resource = channel_pb2.EventarcChannel()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.third_party_provider):
            resource.third_party_provider = Primitive.to_proto(
                self.third_party_provider
            )
        if Primitive.to_proto(self.crypto_key_name):
            resource.crypto_key_name = Primitive.to_proto(self.crypto_key_name)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ChannelStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return channel_pb2.EventarcChannelStateEnum.Value(
            "EventarcChannelStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return channel_pb2.EventarcChannelStateEnum.Name(resource)[
            len("EventarcChannelStateEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
