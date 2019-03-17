import java.util.*;

public class Main {
    public static void main(String[] args) throws Exception {

        Scanner sc = new Scanner(System.in);
        String S = sc.next();
        int n = S.length();
        
        List<Character>list = new ArrayList<>();
        list.add(S.charAt(0));
        
        for(int i = 0; i < n - 1; i++){
            if(S.charAt(i) != S.charAt(i + 1)){
                list.add(S.charAt(i + 1));
            }
        }

        int ans = list.size() - 1;
        System.out.println(ans);
    }
}
