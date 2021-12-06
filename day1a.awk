BEGIN{
 prev=0;
 curr=0;
 fst=0;
 gt7times=0;
 
}

{
if (fst==0) {
prev=$1;
curr=$0;
fst=1;
}else {
prev=curr;
curr=$0;

	if (curr >  prev ) {
	print curr, "increased "
	gt7times++;
	}
}
}
END{
print gt7times
}

## usage awk -f day1a.awk day1inputfile
