package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MortgageInsurance struct {
	Id                                  int       `orm:"column(id);auto"`
	BatchId                             int       `orm:"column(batch_id);null"`
	Gender                              string    `orm:"column(gender);size(8);null"`
	Smoker                              int       `orm:"column(smoker);null"`
	Age                                 int       `orm:"column(age);null"`
	MortgageRecent                      int       `orm:"column(mortgage_recent);null"`
	MortgageAmount                      int       `orm:"column(mortgage_amount);null"`
	LenderName                          string    `orm:"column(lender_name);size(45);null"`
	LenderBranch                        string    `orm:"column(lender_branch);size(45);null"`
	LenderCity                          string    `orm:"column(lender_city);size(45);null"`
	CountryResident                     string    `orm:"column(country_resident);size(45);null"`
	ResidentStatus                      string    `orm:"column(resident_status);size(45);null"`
	Dob                                 time.Time `orm:"column(dob);type(date);null"`
	Height                              int       `orm:"column(height);null"`
	Weight                              int       `orm:"column(weight);null"`
	CancerStatus                        int       `orm:"column(cancer_status);null"`
	CancerTypes                         string    `orm:"column(cancer_types);size(255);null"`
	DiabetesStatus                      int       `orm:"column(diabetes_status);null"`
	BloodDisorderStatus                 int       `orm:"column(blood_disorder_status);null"`
	DiabetiesComplications              int       `orm:"column(diabeties_complications);null"`
	DiabetesComplicationsDetails        string    `orm:"column(diabetes_complications_details);size(255);null"`
	HighBloodPressure                   int       `orm:"column(high_blood_pressure);null"`
	HighCholesterol                     int       `orm:"column(high_cholesterol);null"`
	HeartProblems                       int       `orm:"column(heart_problems);null"`
	HighBloodPressureTreatment          string    `orm:"column(high_blood_pressure_treatment);size(45);null"`
	HighBloodPressureHeartKidneyProblem int       `orm:"column(high_blood_pressure_heart_kidney_problem);null"`
	HighCholesterolTreatment            string    `orm:"column(high_cholesterol_treatment);size(255);null"`
	CholesterolChecked12m               int       `orm:"column(cholesterol_checked_12m);null"`
	LastCholesterolCheckStatus          string    `orm:"column(last_cholesterol_check_status);size(45);null"`
	LastCholesterolCheckReading         string    `orm:"column(last_cholesterol_check_reading);size(45);null"`
	HeartProblemsDetails                string    `orm:"column(heart_problems_details);size(255);null"`
	GastroProblemsStatus                string    `orm:"column(gastro_problems_status);size(255);null"`
	KidneyBladderProblemsStatus         string    `orm:"column(kidney_bladder_problems_status);size(255);null"`
	BreathingLungProblemStatus          string    `orm:"column(breathing_lung_problem_status);size(255);null"`
	GastroProblemsDetails               string    `orm:"column(gastro_problems_details);size(255);null"`
	HerniaRepaired                      int       `orm:"column(hernia_repaired);null"`
	NumGastritis12m                     string    `orm:"column(num_gastritis_12m);size(8);null"`
	GallstonesRemoved                   int       `orm:"column(gallstones_removed);null"`
	GallstonesOngoingIssues             int       `orm:"column(gallstones_ongoing_issues);null"`
	UlcerHospitlisation2d               int       `orm:"column(ulcer_hospitlisation_2d);null"`
	OtherGastroProblems12m              int       `orm:"column(other_gastro_problems_12m);null"`
	NumGastroIssues5y                   int       `orm:"column(num_gastro_issues_5y);null"`
	KidneyProblemDetails                string    `orm:"column(kidney_problem_details);size(255);null"`
	BreathingProblemsDetails            string    `orm:"column(breathing_problems_details);size(255);null"`
	RiskyOccupation                     string    `orm:"column(risky_occupation);size(255);null"`
	RecreationalActivities              string    `orm:"column(recreational_activities);size(255);null"`
	TypeOfCancerWhenDiagnosed           string    `orm:"column(type_of_cancer_when_diagnosed);size(255);null"`
	CancerMedicatedTreated              string    `orm:"column(cancer_medicated_treated);size(255);null"`
	CancerStateRemission                string    `orm:"column(cancer_state_remission);size(255);null"`
	DiabetesFirstDiagnosed              string    `orm:"column(diabetes_first_diagnosed);size(255);null"`
	DiabetesMedicationControl           string    `orm:"column(diabetes_medication_control);size(255);null"`
	DiabetesTests12m                    string    `orm:"column(diabetes_tests_12m);size(255);null"`
	BloodDisorderDiagnosed              string    `orm:"column(blood_disorder_diagnosed);size(255);null"`
	BloodDisorderMedication             string    `orm:"column(blood_disorder_medication);size(255);null"`
	BloodDisorderTests12m               string    `orm:"column(blood_disorder_tests_12m);size(255);null"`
	BloodPressureHighDiagnosed          string    `orm:"column(blood_pressure_high_diagnosed);size(255);null"`
	BloodPressureMedication             string    `orm:"column(blood_pressure_medication);size(255);null"`
	BloodPressureTest                   string    `orm:"column(blood_pressure_test);size(255);null"`
	CholesterolDiagnosed                string    `orm:"column(cholesterol_diagnosed);size(255);null"`
	CholestrolMedication                string    `orm:"column(cholestrol_medication);size(255);null"`
	CholesterolTest                     string    `orm:"column(cholesterol_test);size(255);null"`
	HeartConditionDiagnosed             string    `orm:"column(heart_condition_diagnosed);size(255);null"`
	HeartConditionMedication            string    `orm:"column(heart_condition_medication);size(255);null"`
	HeartConditionTest12m               string    `orm:"column(heart_condition_test_12m);size(255);null"`
	GastroIntestinalDiagnosed           string    `orm:"column(gastro_intestinal_diagnosed);size(255);null"`
	GastroIntestinalMedication          string    `orm:"column(gastro_intestinal_medication);size(255);null"`
	GastroIntestinalCurrentState        string    `orm:"column(gastro_intestinal_current_state);size(255);null"`
	KidneyBladderDiagnosed              string    `orm:"column(kidney_bladder_diagnosed);size(255);null"`
	KidneyBladderMedication             string    `orm:"column(kidney_bladder_medication);size(255);null"`
	KidneyBladderCurrentState           string    `orm:"column(kidney_bladder_current_state);size(255);null"`
	LungProblemDiagnosed                string    `orm:"column(lung_problem_diagnosed);size(255);null"`
	LungProblemMedication               string    `orm:"column(lung_problem_medication);size(255);null"`
	LungProblemCurrentState             string    `orm:"column(lung_problem_current_state);size(255);null"`
	NeurologicalDiagnosed               string    `orm:"column(neurological_diagnosed);size(255);null"`
	NeurologicalMedication              string    `orm:"column(neurological_medication);size(255);null"`
	NeurologicalCurrentState            string    `orm:"column(neurological_current_state);size(255);null"`
	MuscularSkeletalDiagnosed           string    `orm:"column(muscular_skeletal_diagnosed);size(255);null"`
	MuscularSkeletalMedication          string    `orm:"column(muscular_skeletal_medication);size(255);null"`
	MuscularSkeletalCurrentStatus       string    `orm:"column(muscular_skeletal_current_status);size(255);null"`
	MentalHealthDiagnosed               string    `orm:"column(mental_health_diagnosed);size(255);null"`
	MentalHealthCausedEventIllness      string    `orm:"column(mental_health_caused_event_illness);size(255);null"`
	MentalHealthMedication              string    `orm:"column(mental_health_medication);size(255);null"`
	AlcoholTypicalUse                   string    `orm:"column(alcohol_typical_use);size(255);null"`
	AlcoholConsultedProfessional        string    `orm:"column(alcohol_consulted_professional);size(255);null"`
	AlcoholAlcoholismOrCounselling      string    `orm:"column(alcohol_alcoholism_or_counselling);size(255);null"`
	IllegalDrugsUse                     string    `orm:"column(illegal_drugs_use);size(255);null"`
	IllegalDrugsConsultedProfessional   string    `orm:"column(illegal_drugs_consulted_professional);size(255);null"`
	IllegalDrugsHaveStoppedPeriod       string    `orm:"column(illegal_drugs_have_stopped_period);size(255);null"`
	CurrentMedicalAdviceSymptoms        string    `orm:"column(current_medical_advice_symptoms);size(255);null"`
	CurrentMedicalAdviceProfession      string    `orm:"column(current_medical_advice_profession);size(255);null"`
	CurrentMedicalAdviceTests           string    `orm:"column(current_medical_advice_tests);size(255);null"`
	FamilyMemberIllnessDetails          string    `orm:"column(family_member_illness_details);size(255);null"`
	Family2OrMoreHeartDisease           int       `orm:"column(family_2_or_more_heart_disease);null"`
	Family2OrMoreStroke                 int       `orm:"column(family_2_or_more_stroke);null"`
	KidneyDiseasePolycystic             int       `orm:"column(kidney_disease_polycystic);null"`
	Family2OrMoreDiabetes               int       `orm:"column(family_2_or_more_diabetes);null"`
	CancerTypeDiagnosed                 string    `orm:"column(cancer_type_diagnosed);size(255);null"`
	FamilyCountColonRectal60            int       `orm:"column(family_count_colon_rectal_60);null"`
	ColonRectalBefore50                 int       `orm:"column(colon_rectal_before_50);null"`
	FamilyCountProstateCancer60         int       `orm:"column(family_count_prostate_cancer_60);null"`
	ProstateCancerBefore50              int       `orm:"column(prostate_cancer_before_50);null"`
	Family2MoreSameCancer60             int       `orm:"column(family_2_more_same_cancer_60);null"`
	CancerBefore50                      int       `orm:"column(cancer_before_50);null"`
	FamilyCountTesticularCancer60       int       `orm:"column(family_count_testicular_cancer_60);null"`
	TesticularCancerBefore50            int       `orm:"column(testicular_cancer_before_50);null"`
	FamilyCountMs60                     int       `orm:"column(family_count_ms_60);null"`
	FamilyCountParkinsons60             int       `orm:"column(family_count_parkinsons_60);null"`
	FamilyCountMotorNeurone60           int       `orm:"column(family_count_motor_neurone_60);null"`
}

func (t *MortgageInsurance) TableName() string {
	return "mortgage_insurance"
}

func init() {
//	orm.RegisterModel(new(MortgageInsurance))
}

// AddMortgageInsurance insert a new MortgageInsurance into database and returns
// last inserted Id on success.
func AddMortgageInsurance(m *MortgageInsurance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMortgageInsuranceById retrieves MortgageInsurance by Id. Returns error if
// Id doesn't exist
func GetMortgageInsuranceById(id int) (v *MortgageInsurance, err error) {
	o := orm.NewOrm()
	v = &MortgageInsurance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMortgageInsurance retrieves all MortgageInsurance matches certain condition. Returns empty list if
// no records exist
func GetAllMortgageInsurance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MortgageInsurance))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MortgageInsurance
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMortgageInsurance updates MortgageInsurance by Id and returns error if
// the record to be updated doesn't exist
func UpdateMortgageInsuranceById(m *MortgageInsurance) (err error) {
	o := orm.NewOrm()
	v := MortgageInsurance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMortgageInsurance deletes MortgageInsurance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMortgageInsurance(id int) (err error) {
	o := orm.NewOrm()
	v := MortgageInsurance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MortgageInsurance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
