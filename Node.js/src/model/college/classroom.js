export class Classroom {
    students = [];
}

export function sortStudentByGpa() {
    let classroom = JSON.parse(localStorage.getItem("classroom"));
    let students = classroom.students;

    let tempStudent;
    for (let index = students.length - 1; index>0; index--){
        for (let index2 = 0; index2 < index; index2++){
            let currentStudent = students[index];
            let currentGpa = currentStudent.gpa;
            let startStudent = students[index2];
            let startGpa = startStudent.gpa;
            if (currentGpa > startGpa){
                tempStudent = currentStudent;
                students[index] = startStudent;
                students[index2] = tempStudent;
            }
        }
    }
    
    classroom = new Classroom();
    classroom.students = students;
    localStorage.setItem("classroom", JSON.stringify(classroom));
}