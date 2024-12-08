import java.io.File;  // Import the File class
import java.io.FileNotFoundException;  // Import this class to handle errors
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class day1 {
  public static int calculateDistance(ArrayList<Integer> l, ArrayList<Integer> r) {
    Collections.sort(l);
    Collections.sort(r);
    
    int sol = 0;
    for (int i = 0; i < l.size(); i++) {
        sol += Math.abs(l.get(i) - r.get(i));
    }
    return sol;
  }

  public static int caculateSimilarityScore(ArrayList<Integer> l, ArrayList<Integer> r) {
    Map<Integer, Integer> occurrances = new HashMap<>();
    int sol = 0;
    for (int n : r) {
      occurrances.put(n, occurrances.getOrDefault(n, 0) +1);
    }
    for (int n: l) {
      sol += n * occurrances.getOrDefault(n, 0);
    }
    return sol;
  }

  public static void main(String[] args) {
    try {
      File myObj = new File("day1input.txt");
      ArrayList<Integer> l = new ArrayList<>();
      ArrayList<Integer> r = new ArrayList<>();
        try (Scanner myReader = new Scanner(myObj)) {
            while (myReader.hasNextLine()) {
                String data = myReader.nextLine();
                String[] line = data.split("(\\s+)");
                // System.out.println(Arrays.toString(line));
                // System.out.println(data);
                l.add(Integer.valueOf(line[0]));
                r.add(Integer.valueOf(line[1]));
            }
            myReader.close();
            int sol = calculateDistance(l, r);

            System.out.println("Day 1 pt 1 solution" + sol);
            System.out.println("Day 1 pt 2 sol: " + caculateSimilarityScore(l, r));
        }
    } catch (FileNotFoundException e) {
      System.out.println("An error occurred.");
    }
  }
}