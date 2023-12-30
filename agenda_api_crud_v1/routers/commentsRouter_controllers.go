package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:EmailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"] = append(beego.GlobalControllerRouter["github.com/felimarod/ejercicio_crud_mid_api/agenda_api_crud_v1/controllers:NumeroTelefonicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
