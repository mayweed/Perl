//exo1-3
//
#include<stdio.h>

int main()
{
	int fahr, celsius;
	int lower, upper, step;

	lower = 0;
	upper =300;
	step = 20;
	printf ("Table de conversion fahrenheit/Celsius\n");

	fahr = lower;
	while (fahr <= upper){
		celsius = 5 * (fahr-32) / 9;
		printf ("%d\t%d\n", fahr, celsius);
		fahr += step;
	}
	return 0;
}
