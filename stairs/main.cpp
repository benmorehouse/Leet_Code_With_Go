#include <iostream>

void stairs(int stairs){
	if (stairs<=0){
		std::cout<<"invalid"<<std::endl;
	}else{
		int num_stairs = 0;
		while(stairs>1){
			stairs-=2;
			num_stairs++;
		}
		std::cout<<"can do stairs 2 at a time: "<<num_stairs<<std::endl;
		if (num_stairs%2 != 0){
			std::cout<<"can then take the final stairs like this"<<num_stairs%2<<std::endl;
		}
		return;
	}
}

int main(){
	std::cout<<"Hello, how many stairs are you trying to go up"<<std::endl;
	int input;
	std::cin>>input;
	std::cout<<"You entered: "<<input<<std::endl;
	stairs(input);
}

	
