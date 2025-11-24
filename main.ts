/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program asks the user to supply their name, age, and whether they are a student (true/false) and tells what life stage they are in.
 */

// Input
const yourName: string = prompt("What is your name?") || "Anonymous";
const age: number = parseInt(prompt("How old are you?") || "0");
const student: string = prompt("Are you a student? (true/false)") || "false";

// Output
if (student == "true") {
  // Student case
  if (age >= 13 && age <= 19) {
    console.log(`Student ${yourName} is a teenager.`);
  } else if (age >= 5 && age <= 12) {
    console.log(`Student ${yourName} is a child.`);
  } else {
    console.log(`${yourName} is in a different life stage.`);
  }
} else {
  // Not a student case
  if (age >= 20 && age <= 30) {
    console.log(`${yourName} is a young adult.`);
  } else {
    console.log(`${yourName} is in a different life stage.`);
  }
}
