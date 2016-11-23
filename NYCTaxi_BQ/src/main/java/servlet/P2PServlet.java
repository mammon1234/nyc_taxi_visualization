package servlet;

import bigquery.BigQueryClient;
import com.google.api.services.bigquery.model.TableCell;
import com.google.api.services.bigquery.model.TableRow;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@WebServlet(name = "P2PServlet")
public class P2PServlet extends HttpServlet {

    private SimpleDateFormat sfg = new SimpleDateFormat("MM/dd/yyyy hh:mm a");
    private SimpleDateFormat outfg = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");

    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        doGet(request, response);
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        String attr = request.getParameter("attr");
        String stime = request.getParameter("stime"); // 04/18/2016 12:00 AM
        String etime = request.getParameter("etime");
        String area = request.getParameter("area");
        if (attr != null && stime != null && etime != null && area != null) {
            Date st = null;
            Date et = null;
            try {
                st = sfg.parse(stime);
                et = sfg.parse(etime);
            } catch (ParseException e) {
                e.printStackTrace();
            }
            String[] locations = area.split("_");

            String sql = "SELECT\n" + attr +
                    " FROM\n" +
                    "  [NYC.Trips]\n" +
                    "WHERE\n" +
                    "  pickup_latitude > \""+ locations[0] +"\"\n" +
                    "  AND pickup_latitude < \""+ locations[2] +"\"\n" +
                    "  AND pickup_longitude > \""+ locations[3] +"\"\n" +
                    "  AND pickup_longitude < \""+ locations[1] +"\"\n" +
                    "  AND pickup_datetime > \""+ outfg.format(st) +"\"" +
                    "  AND dropoff_datetime < \""+ outfg.format(et) +"\"";

            List<TableRow> rows = BigQueryClient.executeQuery(sql);
            List<Double> result = new ArrayList<Double>();
            Map<String, Object> map = new HashMap<String, Object>();

            if (rows != null) {
                map.put("result", "success");
                for (TableRow row : rows) {
                    for (TableCell field : row.getF()) {
                        String v = (String) field.getV();
                        result.add(Double.valueOf(v));
                    }
                }
            } else {
                map.put("result", "fail");
            }
            map.put("list", result);
            Gson gson = new GsonBuilder().create();
            String json = gson.toJson(map);
            response.setContentType("application/json");
            response.setHeader("Access-Control-Allow-Origin", "*");
            response.setCharacterEncoding("UTF-8");
            response.getWriter().write(json);
        }

    }
}
