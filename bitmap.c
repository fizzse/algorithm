#include <stdio.h>
#include <stdlib.h>

typedef struct BIT
{
	unsigned char b1 : 1;
	unsigned char b2 : 1;
	unsigned char b3 : 1;
	unsigned char b4 : 1;
	unsigned char b5 : 1;
	unsigned char b6 : 1;
	unsigned char b7 : 1;
	unsigned char b8 : 1;
}bit;

int main()
{
	bit* p1 =(bit*) malloc(sizeof(bit) * 4);
	int* p2 = (int*)p1;
	scanf("%d", p2);
	for (int i = 3; i >= 0; i--)
	{
		printf("%d", p1[i].b8);
		printf("%d", p1[i].b7);
		printf("%d", p1[i].b6);
		printf("%d", p1[i].b5);
		printf("%d", p1[i].b4);
		printf("%d", p1[i].b3);
		printf("%d", p1[i].b2);
		printf("%d", p1[i].b1);
		printf("  ");
	}

    printf("\n");
}
