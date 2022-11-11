const LocalStorage = require("node-localstorage").LocalStorage;
global.localStorage = new LocalStorage(`./tempStorage`);

const path = require('path');
const express = require('express');

const router = express.Router();

router.use('/welcome', (req, res) => {
    res.sendFile(path.join(__dirname, '../../assets/html', 'welcome.html'));
});

router.use('/main', (req, res) => {
    res.sendFile(path.join(__dirname, '../../assets/html', 'main.html'));
});

router.use('/add-student', (req, res) => {
    res.sendFile(path.join(__dirname, '../../assets/html', 'addStudent.html'));
});

router.use('/show-students', (req, res) => {
    res.sendFile(path.join(__dirname, '../../assets/html', 'showStudents.html'));
});

router.get('/get-student/:studentId', (req, res) => {
    res.sendFile(path.join(__dirname, '../../assets/html', 'showStudent.html'));
});

router.use('/show-peak', (req, res) => {
    res.sendFile(path.join(__dirname, '../../assets/html', 'showPeak.html'));
});


module.exports = router;