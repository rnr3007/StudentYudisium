if (location.pathname == "/show-students"){
    let backBtn = document.getElementById("backBtn");    
    showStudents();

    backBtn.onclick = function() {
        location.href = `${location.origin}/main`;
    };
} else if (location.pathname == "/show-peak") {
    let backBtn = document.getElementById("backBtn");    
    showPeak();

    backBtn.onclick = function() {
        location.href = `${location.origin}/main`;
    };
} else {
    let backBtn = document.getElementById("backBtn");    
    showStudent();

    backBtn.onclick = function() {
        location.href = `${location.origin}/show-students`;
    };
}

function showStudents() {
    let classroom = JSON.parse(localStorage.getItem("classroom"));
    let students = classroom.students;
    
    if (students != undefined && students.length != 0){
        for (let index = 0; index < students.length; index++){
            document.getElementById("tableStudents").innerHTML += 
                `<tr>` +
                    `<td class="tableCenterCell"><a href="${location.origin}/get-student/${index}">${students[index].name}</a></td>` +
                    `<td class="tableCenterCell">${students[index].gpa}</td>` + 
                `</tr>`
        }
    } else {
        document.getElementById("alertTableStudents").innerHTML =
            `<p class="alertText centerParagraph">No student data exists</p>`;
    }
}

function showStudent() {
    let studentId = location.pathname.replace("/get-student/","") * 1;
    try {
        let classroom = JSON.parse(localStorage.getItem("classroom"));
        let student = classroom.students[studentId];
        let courses = student.courses;
        console.log(courses);

        document.getElementById("studentName").innerHTML = student.name;
        document.getElementById("studentCourseAmount").innerHTML += student.courseAmount;
        for (let index = 0; index < courses.length; index++) {
            document.getElementById("studentcourses").innerHTML +=
                `<p class="subHeadingText">${courses[index].courseName}</p>` +
                `<p>Course Mark: ${courses[index].courseLetterMark}</p>` +
                `<p>Course SKS: ${courses[index].courseSks}</p>`;
        }
        document.getElementById("studentTotalSks").innerHTML += student.totalSks;
        document.getElementById("studentGpa").innerHTML += student.gpa;

    } catch (error) {
        alert("Student id doesn't exist!");
        location.href = `${location.origin}/show-students`;
    }
}

function showPeak() {
    let classroom = JSON.parse(localStorage.getItem("classroom"));
    let students = classroom.students;
    
    if (students != undefined && students.length != 0){
        document.getElementById("highestStudent").innerHTML += students[0].name;
        document.getElementById("lowestStudent").innerHTML += students[students.length - 1].name;

    } else {
        document.getElementById("alertTableStudents").innerHTML =
            `<p class="alertText centerParagraph">No student data exists</p>`;
    }
}