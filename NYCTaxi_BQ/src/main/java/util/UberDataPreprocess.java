package util;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * Created by YangYu on 4/6/16.
 */
public class UberDataPreprocess {

    private final static String rpath = "/Users/yangyu/Downloads/uber-raw-data-apr14.csv";
    private final static String wpath = "/Users/yangyu/Downloads/uber/14-4.csv";

    public static void main(String[] args) {

        String line = null;
        SimpleDateFormat wdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        SimpleDateFormat rdf = new SimpleDateFormat("MM/dd/yy HH:mm:ss");

        try {
            BufferedReader br = new BufferedReader(new FileReader(rpath));
            BufferedWriter bw = new BufferedWriter(new FileWriter(wpath));
            while((line = br.readLine()) != null) {
                String[] tmp = line.split(",");
                Date date = rdf.parse(tmp[0]);
                line = wdf.format(date) + "," + tmp[1] + "," + tmp[2] + "," + tmp[3]+"\n";
                bw.write(line);
            }
            br.close();
            bw.close();
        } catch (java.io.IOException e) {
            e.printStackTrace();
        } catch (ParseException e) {
            e.printStackTrace();
        }
    }

}
