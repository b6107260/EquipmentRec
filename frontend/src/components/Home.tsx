import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบเบิกอุปกรณ์ทางการแพทย์</h1>
        <h4>Requirements</h4>
        <p>
        ระบบเบิกอุปกรณ์ทางการแพทย์ของแผนกผู้ป่วยในที่มีอยู่ในคลัง โดยจะมีรายการอุปกรณ์ทางการแพทย์ อาทิ เครื่องมือทางการแพทย์ เตียง อวัยวะเทียมต่าง ๆ เป็นต้น ซึ่งแต่ละรายการจะมีจำนวน รวมถึงราคาของแต่ละอุปกรณ์ เมื่อแพทย์ login เข้ามาซึ่งใช้งานระบบเบิกอุปกรณ์ทางการแพทย์รายการที่ต้องการ ต้องกรอกข้อมูลเบื้องต้นและข้อมูลใบรับผู้ป่วยใน เมื่อกดบันทึกแล้วข้อมูลจะถูกบันทึกไว้ในบันทึกการเบิกอุปกรณ์ทางการแพทย์ จากนั้นจะถูกเก็บบันทึกไว้ในฐานข้อมูลของระบบเบิกอุปกรณ์ทางการแพทย์
        </p>
        <img src="/img/picture1.jpg" width="900px"></img>
      </Container>
    </div>
  );
}
export default Home;