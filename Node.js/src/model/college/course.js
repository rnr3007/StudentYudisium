export class Course {
    #courseName;
    #courseSks;
    #courseLetterMark;
    
    set setCourseName(courseName) {
        this.courseName = courseName;
    }

    set setCourseSks(courseSks) {
        this.courseSks = courseSks;
    }

    set setCourseLetterMark(courseLetterMark) {
        this.courseLetterMark = courseLetterMark;
    }

    get getCourseName() {
        return this.courseName;
    }

    get getCourseSks() {
        return this.courseSks;
    }

    get getCourseLetterMark() {
        return this.courseLetterMark;
    }

    get getCourseDoubleMark() {
        switch (this.courseLetterMark) {
            case "A":
                return 4.0;
                break;
            case "B":
                return 3.0;
                break;
            case "C":
                return 2.0;
                break;
            case "D":
                return 1.0;
                break;
        }
        return 0.0;
    }
}