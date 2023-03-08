# Given an array A[] of N positive integers which can contain integers from 1 to P where elements can be repeated or can be absent from the array. Your task is to count the frequency of all elements from 1 to N.
arr=["aaa","bbb","ccc","bbb","aaa","aaa"]
d={}
for ch in arr:
	d[ch]=d.get(ch,0)+1
dict(sorted(d.items(), key=lambda item: item[1]))
print(d)
l=[]
for k in d.keys():
	l.append(k)
if len(l)>1:
	print(l[1])
else:
	print(l[0])