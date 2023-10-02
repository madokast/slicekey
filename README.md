# slicekey

可以当作 map key 的 slice

原理是利用反射动态构建一个数组对象，将 slice 元素复制到数组，从而可以作为 key 使用。
