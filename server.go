package main

import (
	"github.com/gin-gonic/gin"
	"go_api/config"
	"go_api/delivery/api"
	"go_api/delivery/middleware"
)

type AppServer interface {
	Run()
}

type productServer struct {
	router *gin.Engine
	config config.Config
}

func (p *productServer) initHandlers() {
	token := p.config.Get("productapp.api.token")
	p.router.Use(middleware.DummyMiddleWare)
	p.router.Use(middleware.ErrorMiddleWare())
	p.router.Use(middleware.TokenAuthMiddleWare(token))
	p.v1()
}

func (p *productServer) v1() {
	productApiGroup := p.router.Group("/product")
	api.NewProductApi(productApiGroup)

	//pingApiGroup := p.router.Group("/ping")
	//api.NewProductApi(productApiGroup)
}

func (p *productServer) Run() {
	p.initHandlers()
	listeningAddr := p.config.Get("productapp.api.url")
	err := p.router.Run(listeningAddr)
	if err != nil {
		panic("Server not starting")
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.New(".", "config")
	return &productServer{
		router: r,
		config: c,
	}
}
