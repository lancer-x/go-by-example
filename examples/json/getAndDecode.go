package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"github.com/tidwall/gjson"
)

type carInfo struct {
	Data []dataItem
	Links map[string]interface{}
	meta map[string]interface{}
}

type dataItem struct {
	Id int `mapstructure:"id" json:"id"`
	CarId int `mapstructure:"car_id" json:"car_id"`
	CarSourceId int `mapstructure:"car_source_id" json:"car_source_id"`
	LicenseYear int `mapstructure:"license_year" json:"license_year"`
	LicenseMonth int `mapstructure:"license_month" json:"license_month"`
	CarImages []CarImageItem `mapstructure:"car_images" json:"car_images"`
	BrightPoint []string `mapstructure:"bright_point" json:"bright_point"`
	Color int `mapstructure:"color" json:"color"`
	Description string `mapstructure:"description" json:"description"`
}

type CarImageItem struct {
	Ind int `mapstructure:"ind" json:"ind"`
	ImageUrl string `mapstructure:"image_url,omitempty" json:"image_url,emitempty"`
	Category string `mapstructure:"category,omitempty" json:"category,emitempty"`
	CategoryDesc string `mapstructure:"category_desc,omitempty" json:"category_desc,emitempty"`
}

const url  = "http://open.guazi.com/api/v1/cars?account=dongchedi&token=732bef746ed7d330d067cff37fe77628a9c590fe"
func main66()  {
	data := getData()
	//parseJsonStr(data)
	direct2Struct(data)
	getProperty(string(data))
}

func getData() []byte {
	return []byte(rawData())
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	retStr,_ := ioutil.ReadAll(resp.Body)
	return retStr
}

func rawData() string {
	return `{"data":[{"id":3013038592,"car_id":173955,"car_source_id":102020056,"license_year":2016,"license_month":8,"car_images":[{"ind":1,"image_url":"https://image1.guazistatic.com/qn20101209595180d6d346e7fc0960dc5a3cf47ff4440e.jpg?imageView2/1/w/620/h/430/q/88"},{"ind":2,"image_url":"https://image1.guazistatic.com/qn201012095951d5bd6b0834c7c8e3b3da7eabbdef40b4.jpg?imageView2/1/w/620/h/430/q/88","category":"正前","category_desc":"前脸完好"},{"ind":3,"image_url":"https://image1.guazistatic.com/qn201012095951c53e1a6c1ff3f83aaef76297d82c67f9.jpg?imageView2/1/w/620/h/430/q/88"},{"ind":4,"image_url":"https://image1.guazistatic.com/qn201012095951c708f0ef0afa4d434617ef03db67af05.jpg?imageView2/1/w/620/h/430/q/88","category":"正侧","category_desc":"漆面保持较好，车身结构无修复，无重大事故"}],"thumb_url":"https://image1.guazistatic.com/qn20101209595180d6d346e7fc0960dc5a3cf47ff4440e.jpg?imageView2/1/w/620/h/430/q/88","transfer_times":0,"price":200000,"guide_price":384039.32,"transfer_fee":0,"road_haul":100700,"car_type":1,"color":2,"description":"经检测，该车排除火烧情况；个别部位有磨损痕迹；关键事故部件中未发现变形异常项；无钣金痕迹；无焊接、切割动作。发动机舱内部无拆卸拆解痕迹；核心部件无更换记录；无漏油痕迹。该车覆盖件少量喷漆修复，少量钣金修复，少量更换，不影响行车安全。车辆外观有轻微划痕刮蹭；详细结果请查看检测报告。","car_source_status":0,"publish_time":"2020-10-18 23:36:54","city_id":26,"phone_400":"4000609028","transfer_num":0,"seller_name":"。先生","seller_description":"","district_name":"余杭区","report_url":"https://m.guazi.com/hz/b898168fcfc690d6x?act=getReportPage","bright_point":["前排座椅加热","倒车雷达","倒车影像系统","多功能方向盘"],"wap_url":"https://m.guazi.com/hz/b898168fcfc690d6x?ca_s=bd_dongchedi&ca_n=donghedimc","pc_url":"https://guazi.com/hz/b898168fcfc690d6x.htm?ca_s=bd_dongchedi&ca_n=donghedimc","product_name":"奔驰C级 2016款 C 200 L 运动型"},{"id":3013038575,"car_id":182425,"car_source_id":102306053,"license_year":2018,"license_month":11,"car_images":[{"ind":1,"image_url":"https://image1.guazistatic.com/qn2010181648515fdf423d310b3d7ed18d2025291ca04f.jpg?imageView2/1/w/620/h/430/q/88"},{"ind":2,"image_url":"https://image1.guazistatic.com/qn201018164851ef250434d767a3e4b0d8c6c655baad88.jpg?imageView2/1/w/620/h/430/q/88","category":"正前","category_desc":"前脸完好"},{"ind":3,"image_url":"https://image1.guazistatic.com/qn201018164851608dd1ef9d52c86db6f47b9d142f3ad2.jpg?imageView2/1/w/620/h/430/q/88"}],"thumb_url":"https://image1.guazistatic.com/qn2010181648515fdf423d310b3d7ed18d2025291ca04f.jpg?imageView2/1/w/620/h/430/q/88","transfer_times":1602777600,"price":595000,"guide_price":877059.83,"transfer_fee":0,"road_haul":38800,"car_type":1,"color":7,"description":"经检测，该车排除火烧、泡水、调表情况；关键事故部件中未发现变形异常项；无钣金痕迹；无焊接、切割动作。发动机舱内部无拆卸拆解痕迹；核心部件无更换记录；无漏油痕迹。车辆外观有轻微划痕刮蹭；详细结果请查看检测报告。","car_source_status":0,"publish_time":"2020-10-18 22:57:11","city_id":75,"phone_400":"4000609028","transfer_num":1,"seller_name":"l先生","seller_description":"","district_name":"闽侯","report_url":"https://m.guazi.com/fz/a83c515d038a9269x?act=getReportPage","bright_point":["前排座椅加热","倒车雷达","倒车影像系统","GPS导航(多媒体)","多功能方向盘"],"wap_url":"https://m.guazi.com/fz/a83c515d038a9269x?ca_s=bd_dongchedi&ca_n=donghedimc","pc_url":"https://guazi.com/fz/a83c515d038a9269x.htm?ca_s=bd_dongchedi&ca_n=donghedimc","product_name":"宝马7系 2018款 730Li 领先型 卓越套装(进口)"}],"links":{"first":"http://open.guazi.com/api/v1/cars?page=1","last":"http://open.guazi.com/api/v1/cars?page=901","prev":null,"next":"http://open.guazi.com/api/v1/cars?page=2"},"meta":{"current_page":1,"from":1,"last_page":901,"path":"http://open.guazi.com/api/v1/cars","per_page":100,"to":100,"total":90028}}`
}

func parseJsonStr(data []byte) {
	var ret carInfo
	var temp map[string]interface{}

	json.Unmarshal(data, &temp)
	mapstructure.Decode(temp, &ret)
	fmt.Printf("%T %v",ret.Data[0].CarImages[0].ImageUrl, ret.Data[0].CarImages[0].ImageUrl)
}

func direct2Struct(data []byte)  {
	var ret carInfo
	json.Unmarshal(data, &ret)
	fmt.Println(ret)
}

func getProperty(data string)  {
	a := gjson.Get(data, "data.0.car_id")
	fmt.Println(a.Int())
}