import { safeAlphabetString } from "../utils/inputValidator.js";
import { Student } from "../model/user/student.js";
import { Course } from "../model/college/course.js";
import { sortStudentByGpa } from "../model/college/classroom.js";

var submitStudentBtn = document.getElementById("submitStudentBtn");
var backBtn = document.getElementById("backBtn");
var formCourses = document.getElementById("formCourses");
var selectCourseAmount = document.getElementById("inputCourseAmount");

var ids = [];
var studentData = [];

changeCourseForm(selectCourseAmount.value);

backBtn.onclick = function() {
    location.href = `${location.origin}/main`;
};

selectCourseAmount.onchange = function() {
    let studentCourseAmount = selectCourseAmount.value;
    changeCourseForm(studentCourseAmount);
};

submitStudentBtn.onclick = function() {
    studentData = [];

    ids = ["inputName", "inputCourseAmount"];
    let courseAmount = selectCourseAmount.value;
    for (let count = 1; count <= courseAmount; count++){
        ids.push(`inputCourseName${count}`);
        ids.push(`inputCourseSks${count}`);
        ids.push(`inputCourseMark${count}`);
    }
    if (validateStudentData()) {
        addStudentData();
    }
};

function changeCourseForm(studentCourseAmount) {
    formCourses.innerHTML = "";
    if (studentCourseAmount != 0){
        for (let count = 1; count <= studentCourseAmount; count++){
            formCourses.innerHTML += 
            `<p class="boldText">Course ${count}</p>` +
            `<label for="inputCourseName${count}">Course name&emsp;&nbsp;</label>` +
            `<input type="text" id="inputCourseName${count}" class="inputShortText" value="">` +
            `<p class="alertText" id="alertInputCourseName${count}"></p>` +
            `<label for="inputCourseSks${count}">Course SKS&emsp;&emsp;</label>` +
            `<select id="inputCourseSks${count}">` +
                `<option value="0">Select value</option>` +
                `<option value="2">2</option>` + 
                `<option value="3">3</option>` +
                `<option value="4">4</option>` +
                `<option value="5">5</option>` +
                `<option value="6">6</option>` +
            `</select>` +
            `<p class="alertText" id="alertInputCourseSks${count}"></p>` +
            `<label for="inputCourseMark${count}">Course mark&emsp;&nbsp;</label>` +
            `<select id="inputCourseMark${count}">` +
                `<option value="0">Select value</option>` +
                `<option value="A">A</option>` +
                `<option value="B">B</option>` +
                `<option value="C">C</option>` +
                `<option value="D">D</option>` +
                `<option value="E">E</option>` +
            `</select>` + 
            `<p class="alertText" id="alertInputCourseMark${count}"></p>`
        }
    }
}

function validateStudentData() {
    let isValid = true;

    let regexIdName = new RegExp("name", "i");

    for (let index = 0; index < ids.length ; index++){
        let id = ids[index];
        let alertId = `alert${id.charAt(0).toUpperCase()+id.slice(1)}`;
        if (regexIdName.test(id)){
            let name = document.getElementById(id).value;
            if (!safeAlphabetString(name)){
                isValid = false;
                document.getElementById(alertId)
                    .innerHTML = "Name doesn't meet the requirement (alphabet only)!";
            } else {
                document.getElementById(alertId).innerHTML = "";
                studentData.push(name);
            }
        } else {
            let value = document.getElementById(id).value;
            if (value == "0"){
                isValid = false;
                document.getElementById(alertId)
                    .innerHTML = "Please select the proper option!";
            } else {
                document.getElementById(alertId).innerHTML = "";
                studentData.push(value);
            }
        }
    }
    return isValid;
}

function addStudentData() {
    let student = new Student();
    student.setName = studentData[0];
    student.setCourseAmount = studentData[1];

    student = addStudentCoursesData(student);
    
    let classroom = JSON.parse(localStorage.getItem("classroom"));
    classroom.students.push(student);

    localStorage.setItem("classroom", JSON.stringify(classroom));

    sortStudentByGpa();

    alert("Student sucessfully added");
}

function addStudentCoursesData(student){
    let sksTotal = 0;
    let gpa = 0.0;
    let courses = [];
    let index = 2;
    
    while (index < studentData.length) {
        let course = new Course();
        course.setCourseName = studentData[index];
        course.setCourseSks = studentData[index + 1] * 1;
        course.setCourseLetterMark = studentData[index + 2];
        courses.push(course);
        index += 3;
        sksTotal += course.getCourseSks;
        gpa += (course.getCourseDoubleMark * course.getCourseSks);
    }

    gpa = (gpa / sksTotal).toFixed(2);
    
    student.setCourses = courses;
    student.setTotalSks = sksTotal;
    student.setGpa = gpa;

    return student;
}