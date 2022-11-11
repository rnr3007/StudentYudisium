export class Student {
    #name;
    #courseAmount;
    #courses;
    #totalSks;
    #gpa;
    
    set setName(name) {
        this.name = name;
    }

    set setCourseAmount(courseAmount) {
        this.courseAmount = courseAmount;
    }

    set setCourses(courses){
        this.courses = courses;
    }

    set setTotalSks(totalSks) {
        this.totalSks = totalSks;
    }

    set setGpa(gpa){
        this.gpa = gpa;
    }

    get getName() {
        return this.name;
    }

    get getCourseAmount() {
        return this.courseAmount;
    }

    get getCourses() {
        return this.courses;
    }

    get getTotalSks() {
        return this.totalSks;
    }

    get getGpa(){
        return this.gpa;
    }

    get getPredicate() {
        for (let index = 0; index < this.courses.length; index++){
			if (this.courses[index].getCourseLetterMark == "E"){
				return "PENDING";
			}
		}
		
		if(this.gpa >= 3.5){
			return "CUM LAUDE";
		}
		else if(this.gpa >= 2.75 && this.gpa < 3.5){
			return "VERY SATISFYING";
		}
		else if(this.gpa >= 2 && this.gpa < 2.75){
			return "SATISFYING";
		}
        return "PENDING";
    }
}