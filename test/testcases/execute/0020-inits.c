
int x = 0;

int y = {2};

int main() {
	if (x) return 1;
	if (y != 2) return 1;
	return 0;
}