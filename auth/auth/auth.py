import os

import dotenv
import grpc
import jwt
import redis

from auth import auth_pb2_grpc
from auth import auth_pb2

from utils.jwt import generate_token, decode_token

dotenv.load_dotenv()

r = redis.Redis(host='localhost', port=6379, decode_responses=True)


class AuthService(auth_pb2_grpc.AuthServiceServicer):
    def Login(self, request, context):
        access_token, refresh_token = generate_token(request.id)
        return auth_pb2.JWTToken(
            access_token=access_token,
            refresh_token=refresh_token
        )

    def Refresh(self, request, context):
        try:
            user_id = decode_token(request.refresh_token)['id']
        except jwt.exceptions.DecodeError:
            context.set_details("REFRESH_INVALID")
            return auth_pb2.JWTToken(
                access_token=None,
                refresh_token=None
            )
        except jwt.exceptions.ExpiredSignatureError:
            context.set_details("REFRESH_EXPIRED")
            return auth_pb2.JWTToken(
                access_token=None,
                refresh_token=None
            )

        if r.get(request.refresh_token):
            context.set_details("")
            return auth_pb2.JWTToken(
                access_token=None,
                refresh_token=None
            )
        access_token, refresh_token = generate_token(user_id)
        r.set(request.refresh_token, user_id)
        return auth_pb2.JWTToken(
            access_token=access_token,
            refresh_token=refresh_token
        )

    def CheckToken(self, request, context):
        if r.get(request.refresh_token):
            context.set_code("")
            return auth_pb2.Payload(id=None)
        try:
            user_id = decode_token(request.access_token)['id']
        except jwt.exceptions.DecodeError:
            context.set_details("ACCESS_INVALID")
            return auth_pb2.Payload(id=None)
        except jwt.exceptions.ExpiredSignatureError:
            context.set_details("ACCESS_EXPIRED")
            return auth_pb2.Payload(id=None)
        except Exception as e:
            print(e)
        return auth_pb2.Payload(id=user_id)
