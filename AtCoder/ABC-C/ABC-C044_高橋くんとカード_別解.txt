import java.util.*;

public class Main {
    public static void main(String[] args) throws Exception {

        Scanner sc = new Scanner(System.in);
        int N = sc.nextInt();
        int A = sc.nextInt();
        int[] x = new int[N];
        for(int i = 0; i < N; i++) x[i] = sc.nextInt();

        long[][][] dp = new long[N+1][N+1][A*N+1];
        dp[0][0][0] = 1;
        
        for(int i = 0; i < N; i++){
            for(int j = 0; j < N; j++){
                for(int k = 0; k <= A * N; k++){
                    dp[i+1][j][k] += dp[i][j][k];
                    if(k >= x[i]){
                        dp[i+1][j+1][k] += dp[i][j][k-x[i]]; 
                    }
                }
            }
        }
        long ans = 0;
        for(int i = 1; i <= N; i++) ans += dp[N][i][i*A];
        System.out.println(ans);
    }
}
