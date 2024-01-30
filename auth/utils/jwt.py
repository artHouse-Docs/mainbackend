import os
from datetime import *
import jwt
import dotenv

dotenv.load_dotenv()


def generate_token(user_id):
    payload = lambda data, exp: {
        'id': data,
        'exp': datetime.utcnow() + exp,
    }

    return [
        jwt.encode(payload(user_id, timedelta(seconds=120)), os.getenv('JWT_SECRET'), algorithm='HS256'),
        jwt.encode(payload(user_id, timedelta(days=10)), os.getenv('JWT_SECRET'), algorithm='HS256'),
    ]


def decode_token(token):
    return jwt.decode(token, os.getenv('JWT_SECRET'), algorithms=['HS256'])
