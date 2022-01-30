/*************************************************************************
  > File Name: main.c
  > Author: fjp
  > Mail: fjp@xxx.com
  > Created Time: Sun Jan 30 12:17:55 2022
  > gcc memleak.c -o Memleak -ldl -g
 ************************************************************************/

#define _GNU_SOURCE //放在所有前面
#include <dlfcn.h> //dlsym

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define MEM_COUNT_LEN 128

//hook
#if 0

typedef void *(*malloc_t)(size_t size);
typedef void (*free_t)(void *p);

malloc_t malloc_f;
free_t free_f;
int enable_malloc_hook = 1;
int enable_free_hook = 1;

static int init_hook(){
	malloc_f = dlsym(RTLD_NEXT, "malloc");
	free_f = dlsym(RTLD_NEXT, "free");
}

void *malloc(size_t size){
	if (enable_malloc_hook){
		enable_malloc_hook =0;

		void *ptr = malloc_f(size);

		void *caller = __builtin_return_address(0);
		char buff[MEM_COUNT_LEN]={0};
		sprintf(buff,"./tmp/%p.mem",ptr);

		FILE *fp = fopen(buff,"w");
		fprintf(fp,"line: [addr2line -fe Memleak -a %p] malloc --> addr: %p, size: %lu\n",caller,ptr,size);
		fflush(fp);

		enable_malloc_hook =1;
		return ptr;
	}else{
		return malloc_f(size);
	}
}

void free(void *p) {
	if (p==NULL) {
		return;
	}
	if(enable_free_hook){
		enable_free_hook = 0;
		char buff[MEM_COUNT_LEN]={0};
		sprintf(buff,"./tmp/%p.mem",p);
		if(unlink(buff)<0){
			printf("double free: %p\n",p);
		}
		free_f(p);
		enable_free_hook = 1;
	}else{ 
		free_f(p);
	}
}

#else
void *malloc_hook(size_t size, const char* file, int lineno){
	void *p = malloc(size);
	char buff[MEM_COUNT_LEN]={0};
	sprintf(buff,"./tmp/%p.mem",p);
	FILE *fp = fopen(buff, "w");
	fprintf(fp, "[%s:%d]malloc --> addr: %p size: %lu\n",file,lineno,p, size);
	fflush(fp);
	return p;
}
void free_hook(void *p, const char* file, int lineno){
	char buff[MEM_COUNT_LEN]={0};
	sprintf(buff,"./tmp/%p.mem",p);
	if(unlink(buff)<0){
		printf("double free: %p\n",p);
	}
	free(p);
}

#define malloc(size) malloc_hook(size, __FILE__, __LINE__)
#define free(p)      free_hook(p, __FILE__, __LINE__)
#endif

int main(){
#if 0
	init_hook();
#endif
	void *p1=malloc(10);
	void *p2=malloc(20);
	free(p1);
	void *p3=malloc(30);
	free(p3);
}
