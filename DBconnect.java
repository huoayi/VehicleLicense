import java.sql.Connection;
import java.sql.SQLException;
public class DBconnect {
    public static Connection getConnection() throws SQLException{
        String Drivername = "com.mysql.jdbc.Driver";
        String Url="jdbc:mysql://localhost:3306/tfs";
        String Username = "root";
        String password = "Lcy654321";
        Connection coon = null;
        try {
            Class.forName(Drivername);
        }catch (ClassNotFoundException e){
            e.printStackTrace();
        }
        coon = java.sql.DriverManager.getConnection(Url,Username,password);
        return coon;
    }
}
