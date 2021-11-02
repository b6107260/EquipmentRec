import { DoctorInterface } from "./IDoctor";
import { AdmissionInterface } from "./IAdmission";
import { EquipmentInterface } from "./IEquipment";


export interface RequisitionInterface {


    ID: string,
    RecTime : Date,
    EquipmentID: number,
    Equipment: EquipmentInterface,
    DoctorID: number,
    Doctor: DoctorInterface,
    AdmissionID: number,
    Admission: AdmissionInterface,
    EquipAmount:number,
    EquipCost:EquipmentInterface,
    
   
   }