arr=[2,3,2,3,5]
P=5
d={}
for ele in arr:
    d[ele]=d.get(ele,0)+1
l=[]
for i in range(P):
    if i+1 not in arr:
        l.append(0)
    else:
        l.append(d[i+1])
print(l)