const btnAddStudent = document.getElementById("btnAddStudent");
const btnShowStudents = document.getElementById("btnShowStudents");
const btnShowPeak = document.getElementById("btnShowPeak");

btnAddStudent.onclick = function() {
    location.href = `${location.origin}/add-student`;
};

btnShowStudents.onclick = function() {
    location.href = `${location.origin}/show-students`;
};

btnShowPeak.onclick = function() {
    location.href = `${location.origin}/show-peak`;
};