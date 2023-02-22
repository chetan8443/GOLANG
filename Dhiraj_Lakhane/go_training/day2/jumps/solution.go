package jumps

func Sol(arr[]int)int{

	var i,jump,n=0,0,len(arr)

	for i < n-1 {
		var val int=arr[i]
		if(val==0){
			return -1
		}
		var max int=0
		var maxind int=-1
		if((i+val)==(n-1)){
			return jump+1
		}

		for j:=i+1;j<n && j<=(i+val);j++{
			if ((j + arr[j]) > max) {
				max = j + arr[j];
				maxind = j;
			}
		}
		i = maxind;
		jump++;
		
	}
	return jump

}