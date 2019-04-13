# Bitmap

## 40亿数存储
	2^32 = 4.2 x 10^10

	4Byte可以表示的数的范围是：0~4.2 x 10^10

	4Byte = 32 bit可以表示32个数是否存在
	40亿个数需要存储的空间大约是2^32 / 8 = 512M

## set bit
	1 << bit = 2^bit

	```cpp
		int arr[500M/4]; //40亿个数
		void setbit(int value) {
			if ((arr[value/32] & 1<<(32-value%32-1) )==0){
				arr[value/32] = 1<<(32-value%32-1);
			}
		}

		value/32: 元素表示位存在第value/32的数组元素里
		arr[value/32]里的数为1<<(31-value%32)
		hint:
		arr[i]中的所有数的特征是，最后value%32的数值相同

		例如：
		value1 = value1/32 + value1%32
		value2 = value2/32 + value2%32
		如果value1和value2都落在arr[i]中，那么value1/32和value2/32相同
		arr[value]中不同的数大小取决于value%32,大小为1<<(31-value%32)
	```