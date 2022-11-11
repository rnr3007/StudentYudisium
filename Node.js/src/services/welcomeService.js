import { Classroom } from "../model/college/classroom.js";

let classroom = new Classroom();
localStorage.setItem("classroom", JSON.stringify(classroom));

const btnContinue = document.getElementById("btnContinue");
const btnBack = document.getElementById("btnBack");

btnContinue.onclick = function() {
    location.href = `${location.origin}/main`;
};

btnBack.onclick = function() {
    location.reload();
};