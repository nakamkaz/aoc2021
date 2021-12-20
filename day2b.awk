BEGIN{
 front=0;
 depth=0;
 aim=0; 
}

{
	if ($1~/forward/) {
	front=front+$2;
	depth=depth+aim*$2;
	}

	if ($1~/down/) {
	# depth=depth+$2;
	aim=aim+$2;
	}

	if ($1~/up/) {
	#depth=depth-$2;	
	aim=aim-$2;
	}
}
END{
print front*depth
}

## usage awk -f file.awk dayinput
