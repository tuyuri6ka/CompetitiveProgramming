import java.util.*;

public class Main {
    public static void main(String[] args) throws Exception {

        Scanner sc = new Scanner(System.in);
        
        int N = sc.nextInt();
        int[] a = new int[N];
        
        for(int i = 0; i < N; i++){
            a[i] = sc.nextInt();
        }

        int min = Integer.MAX_VALUE;
        for(int i = -100; i <= 100; i++){
            int cost = 0;
            for(int j = 0; j < N; j++){
                cost += Math.abs(a[j] - i) * Math.abs(a[j] - i);
            }
            if(min > cost){
                min = cost;
            }
        }
        System.out.println(min);
    }        
}
