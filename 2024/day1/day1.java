import java.io.File;  // Import the File class
import java.io.FileNotFoundException;  // Import this class to handle errors
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Scanner;

public class day1 {
  public static void main(String[] args) {
    try {
      File myObj = new File("day1input.txt");
      ArrayList<Integer> l = new ArrayList<>();
      ArrayList<Integer> r = new ArrayList<>();
      Scanner myReader = new Scanner(myObj);
      while (myReader.hasNextLine()) {
        String data = myReader.nextLine();
        String[] line = data.split("(\\s+)");
        System.out.println(Arrays.toString(line));
        System.out.println(data);
        l.add(Integer.parseInt(line[0]));
        r.add(Integer.parseInt(line[1]));
      }
        Collections.sort(l);
        Collections.sort(r);

        System.out.println(l);
        System.out.println(r);
        int sol = 0;
        for (int i = 0; i < l.size(); i++) {
            sol += Math.abs(l.get(i) - r.get(i));
        }
        System.out.println(sol);
      myReader.close();
    } catch (FileNotFoundException e) {
      System.out.println("An error occurred.");
      e.printStackTrace();
    }
  }
}