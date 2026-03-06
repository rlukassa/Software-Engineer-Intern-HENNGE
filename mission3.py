

import hashlib
import hmac
import struct
import time
import base64
import requests
import json

def generateTOTP(secret, time_step=30, t0=0, digits=10, algorithm=hashlib.sha512):
    t = int((time.time() - t0) / time_step)
    msg = struct.pack(">Q", t)
    key = secret.encode("ascii")
    h = hmac.new(key, msg, algorithm).digest()
    offset = h[-1] & 0x0F
    truncated = struct.unpack(">I", h[offset:offset + 4])[0] & 0x7FFFFFFF
    otp = truncated % (10 ** digits)
    return str(otp).zfill(digits)

def main():
    email = "lukasraja72@gmail.com"
    github_url = "https://gist.github.com/rlukassa/27311314dfa990e8db6687c52597128c"

    payload = {
        "github_url": github_url,
        "contact_email": email,
        "solution_language": "python"
    }

    shared_secret = email + "HENNGECHALLENGE003"

    totp = generateTOTP(shared_secret)
    print(f"TOTP: {totp}")
    url = "https://api.challenge.hennge.com/challenges/003"
    headers = {"Content-Type": "application/json"}

    response = requests.post(
        url,
        data=json.dumps(payload),
        headers=headers,
        auth=(email, totp)
    )

    print(f"Status: {response.status_code}")
    print(f"Response: {response.text}")

if __name__ == "__main__":
    main()



