import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";


import { EquipmentInterface } from "../models/IEquipment";
import { DoctorInterface } from "../models/IDoctor";
import { AdmissionInterface } from "../models/IAdmission";
import { RequisitionInterface } from "../models/IReqRecord";
import { TextField } from "@material-ui/core";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function ReqRecCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [doctors, setDoctors] = useState<DoctorInterface>();
  const [equipments, setEquipments] = useState<EquipmentInterface[]>([]);
  const [admissions, setAdmissions] = useState<AdmissionInterface[]>([]);
  //const [med, setMeds] = useState<MedRecordInterface[]>([]);
  const [reqRecord, setReqRecord] = useState<Partial<RequisitionInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleInputChange = (

    event: React.ChangeEvent<{ id?: string; value: any }>
 
  ) => {
 
    const id = event.target.id as keyof typeof ReqRecCreate;
 
    const { value } = event.target;
 
    setReqRecord({ ...reqRecord, [id]: value });
 

  };
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof reqRecord;
    setReqRecord({
      ...reqRecord,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getDoctor = async () => {
    const uid = Number(localStorage.getItem("uid"));
    fetch(`${apiUrl}/route/GetDoctor/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDoctors(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getEquipment = async () => {
    fetch(`${apiUrl}/route/ListEquipment`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setEquipments(res.data);
        } else {
          console.log("else");
        }
      });
  };
  
  const getAdmission= async () => {
    fetch(`${apiUrl}/route/ListAdmission`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setAdmissions(res.data);
        } else {
          console.log("else");
        }
      });
  };


/*
  const getAdmission = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/playlist/watched/user/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        watchVideo.PlaylistID = res.data.ID
        if (res.data) {
          setPlaylists(res.data);
        } else {
          console.log("else");
        }
      });
  };*/

  useEffect(() => {
    getDoctor();
    getEquipment();
    getAdmission();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };
 

  function submit() {
    let data = {
        DoctorID: convertType(doctors?.ID),
        EquipmentID: convertType(reqRecord.EquipmentID),
         // AdmissionID: convertType(reqRecord.AdmissionID),
        RecTime: selectedDate,
        EquipAmount:  convertType(reqRecord.EquipAmount ?? ""),
        AdmissionID: convertType(reqRecord.AdmissionID),
        EquipCost:convertType(reqRecord.EquipmentID), 

    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/route/CreatRequisition`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
        }
      });
  }

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกการเบิกอุปกรณ์
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>แพทย์ผู้เบิก</p>
              <Select
                native
                disabled
                value={reqRecord.DoctorID}
               // onChange={handleChange}
                /*inputProps={{
                  name: "DoctorID",
                }}*/
              >
                <option aria-label="None" value="">

                  {doctors?.Doctor_name}
                </option>

              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ใบรับเข้าผู้ป่วยใน</p>
              <Select
                native
                value={reqRecord.AdmissionID}
                onChange={handleChange}
                inputProps={{
                  name: "AdmissionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกรายชื่อผู้ป่วยใน
                </option>
               
                {admissions.map((item: AdmissionInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}     {item.PatientName}
                  </option> 
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>อุปกรณ์ทางการแพทย์</p>
              <Select
                native
                value={reqRecord.EquipmentID}
                onChange={handleChange}
                inputProps={{
                  name: "EquipmentID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก
                </option>
                {equipments.map((item: EquipmentInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Equipment_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>อุปกรณ์ทางการแพทย์</p>
              <Select
                native
                value={reqRecord.EquipmentID}
                disabled
                onChange={handleChange}
                inputProps={{
                  name: "EquipmentID",
                }}
              >
                <option aria-label="None" value="">
                  ราคาต่อหน่วย
                </option>
                {equipments.map((item:EquipmentInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Equipment_cost}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>        
          <Grid item xs={6}>
          <FormControl fullWidth variant="outlined">
            <p>จำนวน</p>
            <TextField
              label="กรุณากรอกจำนวน"
              id="EquipAmount"
              name="EquipAmount"
              variant="outlined"
              type="string"
              size="medium"
              value={reqRecord.EquipAmount || ""}
              onChange={handleInputChange}
            />
          </FormControl>
          </Grid> {/*
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>test</p>
              <Select
                native
                value={reqRecord.AdmissionID}
                onChange={handleChange}
                inputProps={{
                  name: "AdmissionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณา test
                </option>
                {admissions.map((item: AdmissionInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.PatientID}  {item.Patient_Name}
                  </option>
                ))}
              </Select>
            </FormControl>
                </Grid>*/}
          {/*    field lock
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เพลย์ลิสต์</p>
              <Select
                native
                value={reqRecord.AdmissionID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "AdmissionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเพลย์ลิสต์
                </option>
                <option value={reqRecord?.ID} key={reqRecord?.ID}>
                  {reqRecord?.Admission?.Patient_Name}
                </option>

                
              </Select>
            </FormControl>
              </Grid> */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="RecTime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label=""
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/requisition_records"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" ,backgroundColor:"#626567" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default ReqRecCreate;