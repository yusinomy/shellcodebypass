f=[]
for i in open('output1.txt'):
    f.append(i.strip())
print('var Passwords = []string'+'{'+'"'+'","'.join(f)+'"'+'}')