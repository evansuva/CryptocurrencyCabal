public class gaussSum {

    static int gaussSum (int m) {
	int sum = 0;
	for (int i = 1; i <= m; i++) {
	    sum = sum + i;
	}
	return sum;
    }

    public static void main(String [] args) {
	int x = Integer.MAX_VALUE;
	System.out.println("IntMax: " + x);
	System.out.println("max gaussSum: " + gaussSum(x));
    }
}
