package main

import (
	"io/ioutil"
	"time"
	"testing"
	"os"

)

func Test_createmsg(t *testing.T) {	
	logmsg := createmsg("External","TRF","Request","{test writelog}")
	if logmsg != "External\tTRF\tRequest\t{test writelog}" {
		t.Fatalf("Expected msg to == External\tTRF\tRequest\t{test writelog}\n, but got %s",
			logmsg)
	}

}

func Test_getfilename(t *testing.T) {
	today := time.Now()
	todayString := today.Format("2006-01-02")	
	
	expected := todayString[:4] + todayString[5:7] + todayString[8:10] + ".log"
	
	result := getfilename()

	if result != expected {
		t.Fatalf("Expected msg to == %s, but got %s", expected,result)
	}

}


func Test_writeln(t *testing.T) {
	
	os.Remove("test.log")
	
	logmsg := "External\tTRF\tRequest\t{test writelog}"
	
	writeln("test.log",logmsg)
	
 	_ , err := os.Stat("test.log")
        if os.IsNotExist(err) {
			
			t.Fatalf("File does not exist.")
        }
	
	
	
    b,err := ioutil.ReadFile("test.log") // just pass the file name
    // if err != nil {
    //     fmt.Print(err)
	// }
	
		
	result1 := string(b[:len(b)])
	expected1 := "External\tTRF\tRequest\t{test writelog}\n"

	if result1 != expected1 {
		t.Fatalf("Expected msg1 to == %s, but got %s", expected1,result1)
	}

	logmsg1 :="External\tTRF\tRequest\t{test writelog1}"
	writeln("test.log",logmsg1)

	c, err := ioutil.ReadFile("test.log") // just pass the file name

			
	result2 := string(c[:len(c)])
	expected2 := "External\tTRF\tRequest\t{test writelog}\nExternal\tTRF\tRequest\t{test writelog1}\n"

	if result2 != expected2 {
		t.Fatalf("Expected msg2 to == %s, but got %s", expected2,result2)
	}
}

func Test_addlog(t *testing.T) {
	

	today := time.Now()
	todayString := today.Format("2006-01-02")	
	
	filename := todayString[:4] + todayString[5:7] + todayString[8:10] + ".log"

	os.Remove(filename)	
	
	Addlog("External","TRF","Request","{test writelog}")

 	_ , err := os.Stat(filename)
    if err != nil {
        if os.IsNotExist(err) {
			
			t.Fatalf("File does not exist.")
        }
	
	
	b, err := ioutil.ReadFile(filename) // just pass the file name
    // if err != nil {
    //     fmt.Print(err)
	// }
	
		
	result1 := string(b[27:len(b)])
	expected1 := "External\tTRF\tRequest\t{test writelog}\n"

	if result1 != expected1 {
		t.Fatalf("Expected msg1 to == %s, but got %s", expected1,result1)
	}

	Addlog("External","TRF","Request","{test writelog1}")	

	c, err := ioutil.ReadFile(filename) // just pass the file name

	result2 := string(c[27:64])		
	result3 := string(c[91:len(c)])
	expected2 := "External\tTRF\tRequest\t{test writelog1}\n"

	if result2 != expected1  {
		t.Fatalf("Expected msg1 to == %s, but got %s", expected1,result2)
	}

	if result3 != expected2  {
		t.Fatalf("Expected msg2 to == %s, but got %s", expected2,result3)
	}
}