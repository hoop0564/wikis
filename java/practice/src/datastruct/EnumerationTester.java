package datastruct;

import java.util.Arrays;
import java.util.Enumeration;
import java.util.Vector;

/**
 * @author gzc
 * @version 1.0.0
 */
public class EnumerationTester {
    public static void main(String[] args) {
        Enumeration<String> days;
        Vector<String> dayNames = new Vector<String>();
        dayNames.add("Monday");
        dayNames.add("Tuesday");
        dayNames.add("Wednesday");
        days = dayNames.elements();
        while (days.hasMoreElements()) {
            System.out.println(days.nextElement());
        }
   }
}