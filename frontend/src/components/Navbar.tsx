import React,{useState,useEffect} from "react";
import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import { Link } from "react-router-dom";
import { Button,Drawer } from "@material-ui/core";

import clsx from "clsx"; 
import HomeIcon from "@material-ui/icons/Home";
import AccountCircleIcon from "@material-ui/icons/AccountCircle";
//import MedicationIcon from '@mui/icons-material/Medication';
//import NoteAltIcon from '@mui/icons-material/NoteAlt';
import InsertDriveFileIcon from '@mui/icons-material/InsertDriveFile';
import CreateIcon from '@mui/icons-material/Create';
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";

import Divider from "@material-ui/core/Divider";
import { List } from "@material-ui/core";
import { DoctorInterface } from "../models/IDoctor";

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>

  createStyles({
    root: {
      display: "flex",
    },
    nav : {background : "#626567"},
    title: {
      flexGrow: 1,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    hide: {
      display: "none",
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
      whiteSpace: "nowrap",
    },
    drawerOpen: {
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerClose: {
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      overflowX: "hidden",
      width: theme.spacing(7) + 1,
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9) + 1,
      },
    },
    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },
    a: {
      textDecoration: "none",
      color: "inherit",
    },
  })
);


function Navbar() {
  
  
 const classes = useStyles();
 const [open, setOpen] = React.useState(false);
 const apiUrl = "http://localhost:8080";
 const requestOptions = {
   method: "GET",
   headers: {
     Authorization: `Bearer ${localStorage.getItem("token")}`,
     "Content-Type": "application/json",
   },
 };
  //const [token, setToken] = React.useState<String>("");
  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const menu = [
    { name: "หน้าแรก", icon: <HomeIcon />, path: "/" },
    { name: "เภสัชกร", icon: <AccountCircleIcon />, path: "/doctors" },
    { name: "บันทึกการเบิกอุปกรณ์", icon: <InsertDriveFileIcon/>, path: "/requisition_records" },
    { name: "สร้างบันทึกการเบิกอุปกรณ์", icon: <CreateIcon />, path: "/requisition_record/create" },
  ];

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };
  const theme = useTheme();
  
  const [doctor, setDoctors] = useState<DoctorInterface>();
 
  const getDoctor = async () => {
    const uid = Number(localStorage.getItem("uid"));
    fetch(`${apiUrl}/doctor/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDoctors(res.data);
        } else {
          console.log("else");
        }
      });
  };
  
    useEffect(() => {
      getDoctor();
  }, []);


 return (
  
    <>
      <AppBar
        position="fixed"
        className={clsx(classes.appBar,classes.nav, {
          [classes.appBarShift]: open,
        })}
      >
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            edge="start"
            className={clsx(classes.menuButton, {
              [classes.hide]: open,
            })}
          >
            <MenuIcon />
          </IconButton>
          <img src="/img/med-jar.png" width="50px"></img>
          <Typography variant="h6" className={classes.title}>
            ระบบบันทึกการเบิกอุปกรณ์ทางการแพทย์
          </Typography>
          <Typography variant="subtitle1" >
            {doctor?.Doctor_name} &nbsp;&nbsp;       
          </Typography>
          <Button /*color="inherit"*/ 
          style={{backgroundColor:"#E5E7E9"}}
          onClick={signout}>
            ออกจากระบบ
          </Button>
        </Toolbar>
      </AppBar>
      <Drawer
        variant="permanent"
        className={clsx(classes.drawer, {
          [classes.drawerOpen]: open,
          [classes.drawerClose]: !open,
        })}
        classes={{
          paper: clsx({
            [classes.drawerOpen]: open,
            [classes.drawerClose]: !open,
          }),
        }}
      >
        <div className={classes.toolbar}>
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === "rtl" ? (
              <ChevronRightIcon />
            ) : (
              <ChevronLeftIcon />
            )}
          </IconButton>
        </div>
        <Divider />
        <List>
          {menu.map((item, index) => (
            <Link to={item.path} key={item.name} className={classes.a}>
              <ListItem button>
                <ListItemIcon>{item.icon}</ListItemIcon>
                <ListItemText primary={item.name} />
              </ListItem>
            </Link>
          ))}
        </List>
      </Drawer>
    </>
  )

          }

export default Navbar;