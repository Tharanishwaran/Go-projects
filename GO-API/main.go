package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{

   ID string     
   Title string
   Author string
   Quantity int


}