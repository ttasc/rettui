package ui

import (
	"github.com/rivo/tview"
)

type mainView struct {
    ReqSide  *reqSide
    ResSide  *resSide
}

type reqSide struct {
    *tview.Flex
    Params      *tview.TextArea
    Headers     *tview.TextArea
    BodyTypes   *tview.DropDown
    Body        *tview.TextArea
}

type resSide struct {
    *tview.Pages
    Headers     *tview.Table
    Body        *tview.TextView
}

func newMainView() *mainView {
    mainView := &mainView{
        ReqSide:  newRequestSide(),
        ResSide: newResponseSide(),
    }
    return mainView
}

func (m *mainView) addTo(layout *tview.Grid) {
    layout.AddItem(m.ReqSide, 1, 0, 1, 1, 0, 0, true)
    layout.AddItem(m.ResSide, 1, 1, 1, 1, 0, 0, true)
}

func newRequestSide() *reqSide {
    requestSide := &reqSide{
        Flex:        tview.NewFlex().SetDirection(tview.FlexRow),
        Params:      tview.NewTextArea().SetPlaceholder("URL Params...").SetWordWrap(true),
        Headers:     tview.NewTextArea().SetPlaceholder("<key>: <value>\nExample:\n...\nCache-Control: no-cache\nServer: Microsoft-IIS/10.0\nContent-Type: application/json\n...").SetWordWrap(false),
        BodyTypes:   tview.NewDropDown().SetLabel("Body-type ").SetOptions([]string{"none", "form-data", "x-www-form-urlencoded", "raw"}, nil).SetCurrentOption(0),
        Body:        tview.NewTextArea().SetPlaceholder("Insert data here...").SetWordWrap(true),
    }

    requestSide.Params.SetBorder(true).SetTitle("Request Params").SetTitleAlign(tview.AlignLeft)
    requestSide.Headers.SetBorder(true).SetTitle("Request Headers").SetTitleAlign(tview.AlignLeft)
    requestSide.Body.SetBorder(true).SetTitle("Request Body").SetTitleAlign(tview.AlignLeft)

    requestSide.AddItem(requestSide.Params, 4, 1, false)
    requestSide.AddItem(requestSide.Headers, 10, 1, false)
    requestSide.AddItem(requestSide.BodyTypes, 1, 1, false)
    requestSide.AddItem(requestSide.Body, 0, 1, false)
    return requestSide
}

func newResponseSide() *resSide {
    responseSide := &resSide{
        Pages:       tview.NewPages(),
        Headers:     tview.NewTable(),
        Body:        tview.NewTextView().SetDynamicColors(true).SetWrap(false).SetTextColor(Colors["red"]),
    }

    responseSide.Body.SetText(responseTestText)

    responseSide.SetBorder(true).SetTitle("Response - Ctrl+H/Ctrl+B to switch Headers/Body").SetTitleAlign(tview.AlignLeft)
    responseSide.AddPage("Headers", responseSide.Headers, true, false)
    responseSide.AddPage("Body", responseSide.Body, true, true)
    return responseSide
}

var responseTestText = `{
    "id": 276614493,
    "master_id": 276614493,
    "sku": "8837015340131",
    "name": "Điện Thoại Samsung Galaxy A16 LTE (4GB/128GB) - Đã kích hoạt bảo hành điện tử - Hàng Chính Hãng",
    "url_key": "dien-thoai-samsung-galaxy-a16-8gb-128gb-da-kich-hoat-bao-hanh-dien-tu-hang-chinh-hang-p276614493",
    "url_path": "dien-thoai-samsung-galaxy-a16-8gb-128gb-da-kich-hoat-bao-hanh-dien-tu-hang-chinh-hang-p276614493.html?spid=276615090",
    "short_url": "https://tiki.vn/product-p276614493.html?spid=276615090",
    "type": "configurable",
    "book_cover": null,
    "short_description": "Samsung Galaxy A16 4GB/128GB là chiếc smartphone vừa đẹp vừa mạnh, phù hợp cho mọi nhu cầu từ giải trí đến công việc. Với thiết kế nhỏ gọn, hiện đại, màn hình lớn, camera rõ nét và pin bền, Galaxy A16...",
    "price": 3110000,
    "list_price": 5090000,
    "original_price": 5090000,
    "badges": [
        {
            "code": "new_pdp",
            "text": "v1"
        }
    ],
    "badges_new": [
        {
            "placement": "left_brand",
            "code": "is_authentic",
            "type": "is_authentic",
            "index": 4,
            "icon": "https://salt.tikicdn.com/ts/upload/d7/56/04/b93b8c666e13f49971483596ef14800f.png",
            "icon_width": 89,
            "icon_height": 20,
            "text_color": null,
            "background_color": null,
            "href": null,
            "text": null
        },
        {
            "placement": "right_authentic",
            "code": "return_policy",
            "type": "return_policy",
            "index": 3,
            "icon": "https://salt.tikicdn.com/ts/ta/b1/3f/4e/cc3d0a2dd751a7b06dd97d868d6afa56.png",
            "icon_width": 114,
            "icon_height": 20,
            "text_color": null,
            "background_color": null,
            "href": null,
            "text": null
        }
    ],
    "badges_v3": [
        {
            "placement": "pdp_badge",
            "code": "return_policy",
            "type": "pdp_icon_badge",
            "index": 3,
            "icon": "https://salt.tikicdn.com/ts/ta/b1/3f/4e/cc3d0a2dd751a7b06dd97d868d6afa56.png",
            "icon_width": 114,
            "icon_height": 20,
            "text_color": null,
            "background_color": null,
            "href": null,
            "text": null
        },
        {
            "placement": "pdp_badge",
            "code": "is_authentic",
            "type": "pdp_icon_badge",
            "index": 4,
            "icon": "https://salt.tikicdn.com/ts/upload/d7/56/04/b93b8c666e13f49971483596ef14800f.png",
            "icon_width": 89,
            "icon_height": 20,
            "text_color": null,
            "background_color": null,
            "href": null,
            "text": null
        }
    ],
    "tracking_info": {
        "amplitude": {
            "is_authentic": true,
            "is_freeship_xtra": false,
            "is_hero": false,
            "is_top_brand": true,
            "return_reason": "any_reason"
        }
    },
    "discount": 1980000,
    "discount_rate": 39,
    "rating_average": 4.8,
    "review_count": 44,
    "review_text": "(44)",
    "favourite_count": 0,
    "thumbnail_url": "https://salt.tikicdn.com/cache/280x280/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png",
    "has_ebook": false,
    "inventory_status": "available",
    "inventory_type": "instock",
    "productset_group_name": "",
    "is_fresh": false,
    "seller": null,
    "is_flower": false,
    "has_buynow": true,
    "is_gift_card": false,
    "salable_type": null,
    "data_version": 3300,
    "day_ago_created": 280,
    "all_time_quantity_sold": 497,
    "meta_title": "",
    "meta_description": "",
    "meta_keywords": "",
    "is_baby_milk": false,
    "is_acoholic_drink": false,
    "description": "<h3>Samsung Galaxy A16 4GB/128GB là chiếc smartphone vừa đẹp vừa mạnh, phù hợp cho mọi nhu cầu từ giải trí đến công việc. Với thiết kế nhỏ gọn, hiện đại, màn hình lớn, camera rõ nét và pin bền, Galaxy A16 sẽ luôn bên bạn trong các hoạt động hằng ngày, đáp ứng tốt cả khi bạn cần giải trí hay làm việc.</h3>\n<h3>Màn hình lớn, trải nghiệm rõ nét</h3>\n<p>Samsung Galaxy A16 là lựa chọn phù hợp cho những ai yêu thích giải trí và công nghệ. Với màn hình Super AMOLED 6.7 inch độ phân giải Full HD+, Galaxy A16 mang lại hình ảnh rõ nét và chi tiết. Độ sáng và độ tương phản cao của màn hình giúp bạn thưởng thức mọi nội dung một cách trọn vẹn, dù là xem phim, chơi game hay lướt mạng xã hội.</p>\n<p><a class=\"preventdefault\" href=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024348-476.jpg\"><img class=\" lazyloaded\" title=\"Samsung Galaxy A16 8GB/128GB - Màn hình\" src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024348-476.jpg\" alt=\"Samsung Galaxy A16 8GB/128GB - Màn hình\" data-src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024348-476.jpg\" /></a></p>\n<h3>Thiết kế gọn gàng, trẻ trung và hiện đại</h3>\n<p>Galaxy A16 vừa là chiếc điện thoại tiện ích, vừa là phụ kiện thời trang đồng hành cùng bạn. Thiết kế mỏng nhẹ với các góc bo cong mang lại cảm giác thoải mái khi cầm. Mặt lưng tối giản và hiện đại, camera bố trí riêng lẻ tạo nên nét tinh tế. Sản phẩm có ba màu sắc để bạn lựa chọn: Đen, xanh và xám.</p>\n<p><a class=\"preventdefault\" href=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024349-328.jpg\"><img class=\" lazyloaded\" title=\"Samsung Galaxy A16 8GB/128GB - Màu sắc\" src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024349-328.jpg\" alt=\"Samsung Galaxy A16 8GB/128GB - Màu sắc\" data-src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024349-328.jpg\" /></a></p>\n<h3>Ba camera ghi lại mọi khoảnh khắc trọn vẹn</h3>\n<p>Hệ thống camera của Galaxy A16 nổi bật với khả năng chụp ảnh ấn tượng trong mọi điều kiện. Camera chính 50 MP cho hình ảnh sắc nét và sống động. Camera góc siêu rộng 5 MP giúp mở rộng khung hình, phù hợp cho ảnh phong cảnh hay ảnh nhóm. Camera macro 2 MP cho phép chụp cận cảnh chi tiết. Đồng thời, camera selfie 13 MP giúp bạn tự tin với những bức ảnh chân dung rạng rỡ.</p>\n<p><a class=\"preventdefault\" href=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024344-529.jpg\"><img class=\" lazyloaded\" title=\"Samsung Galaxy A16 8GB/128GB - Hình ảnh camera\" src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024344-529.jpg\" alt=\"Samsung Galaxy A16 8GB/128GB - Hình ảnh camera\" width=\"750\" height=\"419\" data-src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024344-529.jpg\" /></a></p>\n<h3>Hoạt động ổn định, đáp ứng mọi yêu cầu</h3>\n<p>Chiếc điện thoại Samsung này được trang bị vi xử lý MediaTek Helio G99 mạnh mẽ, kết hợp với RAM 4 GB, mang lại hiệu suất vượt trội trong mọi tác vụ. Dù bạn chơi game, xem phim hay thực hiện công việc đa nhiệm, thiết bị luôn vận hành mượt mà, nhanh chóng. RAM 4 GB giúp việc chuyển đổi giữa các ứng dụng trở nên linh hoạt, giữ cho trải nghiệm luôn ổn định và hiệu quả.</p>\n<p><a class=\"preventdefault\" href=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024347-608.jpg\"><img class=\" lazyloaded\" title=\"Samsung Galaxy A16 8GB/128GB - Trải nghiệm game\" src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024347-608.jpg\" alt=\"Samsung Galaxy A16 8GB/128GB - Trải nghiệm game\" data-src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024347-608.jpg\" /></a></p>\n<h3>Bền bỉ, sẵn sàng mọi thử thách</h3>\n<p>Galaxy A16 được trang bị khả năng kháng nước, kháng bụi đạt chuẩn IP54, mang đến sự bảo vệ đáng tin cậy. Nhờ đó, bạn có thể thoải mái sử dụng máy trong điều kiện thời tiết xấu, từ những cơn mưa nhẹ đến môi trường đầy bụi bẩn.</p>\n<h3>Luôn luôn cập nhật, trải nghiệm hoàn hảo</h3>\n<p>Chiếc điện thoại Android này nổibật với cam kết hỗ trợ lâu dài, mang đến 6 lần nâng cấp hệ điều hành và 6 năm cập nhật bảo mật. Đây không chỉ là một thiết bị công nghệ, mà còn là một lựa chọn bền vững, giúp bạn luôn tận hưởng những tính năng mới và sự an toàn tối ưu.</p>\n<p><a class=\"preventdefault\" href=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024346-730.jpg\"><img class=\" lazyloaded\" title=\"Samsung Galaxy A16 8GB/128GB - Cập nhật dài lâu\" src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024346-730.jpg\" alt=\"Samsung Galaxy A16 8GB/128GB - Cập nhật dài lâu\" data-src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024346-730.jpg\" /></a></p>\n<h3>Pin khỏe, sạc nhanh, trải nghiệm liền mạch</h3>\n<p>Galaxy A16 được trang bị viên pin 5000 mAh, cho phép bạn thoải mái sử dụng cả ngày dài chỉ với một lần sạc. Với công nghệ sạc nhanh 25W, bạn có thể nhanh chóng nạp đầy năng lượng và tiếp tục công việc hay giải trí mà không bị gián đoạn.</p>\n<p><a class=\"preventdefault\" href=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024350-351.jpg\"><img class=\" lazyloaded\" title=\"Samsung Galaxy A16 8GB/128GB - Pin, sạc\" src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024350-351.jpg\" alt=\"Samsung Galaxy A16 8GB/128GB - Pin, sạc\" data-src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024350-351.jpg\" /></a></p>\n<h3>Bảo mật vững chắc, riêng tư tuyệt đối</h3>\n<p>Samsung luôn chú trọng bảo mật với công nghệ Knox Vault trên Galaxy A16. Dữ liệu cá nhân như mã PIN, mật khẩu hay thông tin quan trọng của bạn được bảo vệ an toàn, giống như có một lớp \"két sắt\" giúp giữ mọi thứ luôn riêng tư.</p>\n<p><a class=\"preventdefault\" href=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024345-704.jpg\"><img class=\" lazyloaded\" title=\"Samsung Galaxy A16 8GB/128GB - Bảo mật\" src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024345-704.jpg\" alt=\"Samsung Galaxy A16 8GB/128GB - Bảo mật\" data-src=\"https://cdnv2.tgdd.vn/mwg-static/tgdd/Products/Images/42/331207/samsung-galaxy-a16-8gb-128gb-111124-024345-704.jpg\" /></a></p>\n<p>Với thiết kế tinh tế, camera chất lượng, hiệu năng vượt trội và bảo mật tối ưu, Galaxy A16 là lựa chọn hoàn hảo đáp ứng mọi nhu cầu. Đây không chỉ là một chiếc <a title=\"Tham khảo điện thoại kinh doanh tại Thegioididong.com\" href=\"https://www.thegioididong.com/\" target=\"_blank\" rel=\"noopener\">điện thoại</a> đẹp mà còn thông minh, bền bỉ và đáng tin cậy trong mọi tình huống</p><p>Giá sản phẩm trên Tiki đã bao gồm thuế theo luật hiện hành. Bên cạnh đó, tuỳ vào loại sản phẩm, hình thức và địa chỉ giao hàng mà có thể phát sinh thêm chi phí khác như phí vận chuyển, phụ phí hàng cồng kềnh, thuế nhập khẩu (đối với đơn hàng giao từ nước ngoài có giá trị trên 1 triệu đồng).....</p>",
    "images": [
        {
            "base_url": "https://salt.tikicdn.com/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png",
            "is_gallery": true,
            "label": null,
            "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png",
            "medium_url": "https://salt.tikicdn.com/cache/w300/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png",
            "position": null,
            "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png",
            "thumbnail_url": "https://salt.tikicdn.com/cache/200x280/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png"
        },
        {
            "base_url": "https://salt.tikicdn.com/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg",
            "is_gallery": true,
            "label": null,
            "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg",
            "medium_url": "https://salt.tikicdn.com/cache/w300/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg",
            "position": null,
            "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg",
            "thumbnail_url": "https://salt.tikicdn.com/cache/200x280/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg"
        }
    ],
    "warranty_policy": null,
    "brand": {
        "id": 18802,
        "name": "Samsung",
        "slug": "samsung"
    },
    "current_seller": {
        "id": 1,
        "sku": "2158312495847",
        "name": "Tiki Trading",
        "link": "https://tiki.vn/cua-hang/tiki-trading",
        "logo": "https://vcdn.tikicdn.com/ts/seller/d1/3f/ae/13ce3d83ab6b6c5e77e6377ad61dc4a5.jpg",
        "price": 3110000,
        "product_id": "276615090",
        "store_id": 40395,
        "is_best_store": false,
        "is_offline_installment_supported": null
    },
    "other_sellers": [],
    "specifications": [
        {
            "name": "Content",
            "attributes": [
                {
                    "code": "battery_capacity",
                    "name": "Dung lượng pin",
                    "value": "5000 mAh"
                },
                {
                    "code": "bluetooth",
                    "name": "Bluetooth",
                    "value": "v5.3"
                },
                {
                    "code": "brand",
                    "name": "Thương hiệu",
                    "value": "Samsung"
                },
                {
                    "code": "brand_country",
                    "name": "Xuất xứ thương hiệu",
                    "value": "Hàn Quốc"
                },
                {
                    "code": "camera_sau",
                    "name": "Camera sau",
                    "value": "Chính 50 MP &amp; Phụ 5 MP, 2 MP"
                },
                {
                    "code": "camera_truoc",
                    "name": "Camera trước",
                    "value": "13 MP"
                },
                {
                    "code": "cart_slot",
                    "name": "Hỗ trợ thẻ nhớ ngoài",
                    "value": "MicroSD"
                },
                {
                    "code": "chip_do_hoa",
                    "name": "Chip đồ họa (GPU)",
                    "value": "Mali-G57"
                },
                {
                    "code": "chip_set",
                    "name": "Chip set",
                    "value": "MediaTek Helio G99"
                },
                {
                    "code": "connect_khac",
                    "name": "Kết nối khác",
                    "value": "NFC"
                },
                {
                    "code": "cpu_speed",
                    "name": "Tốc độ CPU",
                    "value": "2.2 GHz"
                },
                {
                    "code": "display_type",
                    "name": "Loại/ Công nghệ màn hình",
                    "value": "Super AMOLED"
                },
                {
                    "code": "ho_tro_5g",
                    "name": "Hỗ trợ 5G",
                    "value": "Không"
                },
                {
                    "code": "item_model_number",
                    "name": "Model",
                    "value": "A16"
                },
                {
                    "code": "jack_headphone",
                    "name": "Jack tai nghe",
                    "value": "Type-C"
                },
                {
                    "code": "khe_sim",
                    "name": "Số sim",
                    "value": "2"
                },
                {
                    "code": "loai_pin",
                    "name": "Loại pin",
                    "value": "Li-Ion"
                },
                {
                    "code": "loai_sim",
                    "name": "Loại Sim",
                    "value": "Nano SIM"
                },
                {
                    "code": "origin",
                    "name": "Xuất xứ (Made in)",
                    "value": "Việt Nam"
                },
                {
                    "code": "port_sac",
                    "name": "Cổng sạc",
                    "value": "Type-C"
                },
                {
                    "code": "ram",
                    "name": "RAM",
                    "value": "4GB"
                },
                {
                    "code": "resolution",
                    "name": "Độ phân giải",
                    "value": "Full HD+ (1080 x 2340 Pixels)"
                },
                {
                    "code": "rom",
                    "name": "ROM",
                    "value": "128GB"
                },
                {
                    "code": "screen_size",
                    "name": "Kích thước màn hình",
                    "value": "6.7 inch"
                },
                {
                    "code": "the_ngoai_toi_da",
                    "name": "Hỗ trợ thẻ tối đa",
                    "value": "1 TB"
                },
                {
                    "code": "wifi",
                    "name": "Wifi",
                    "value": "Dual-band (2.4 GHz/5 GHz)"
                }
            ]
        },
        {
            "name": "Operation",
            "attributes": [
                {
                    "code": "is_warranty_applied",
                    "name": "Sản phẩm có được bảo hành không?",
                    "value": "Có"
                },
                {
                    "code": "vat_taxable",
                    "name": "Có thuế VAT",
                    "value": "Có"
                },
                {
                    "code": "warranty_form",
                    "name": "Hình thức bảo hành",
                    "value": "Điện tử"
                },
                {
                    "code": "warranty_time_period",
                    "name": "Thời gian bảo hành",
                    "value": "12 Tháng"
                }
            ]
        }
    ],
    "product_links": [],
    "gift_item_title": "0 quà tặng kèm",
    "configurable_options": [
        {
            "code": "option1",
            "name": "Màu",
            "position": 0,
            "show_preview_image": true,
            "values": [
                {
                    "label": "Xanh Lá"
                },
                {
                    "label": "Xám Bạc"
                },
                {
                    "label": "Đen"
                }
            ]
        }
    ],
    "configurable_products": [
        {
            "child_id": 276614495,
            "discount_rate": 39,
            "id": 276615090,
            "images": [
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png"
                },
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/da/af/09/3da6b36682deb00818b404272bd8be58.jpg"
                }
            ],
            "inventory_status": "available",
            "name": "Điện Thoại Samsung Galaxy A16 LTE (4GB/128GB) - Đã kích hoạt bảo hành điện tử - Hàng Chính Hãng - Xanh Lá",
            "option1": "Xanh Lá",
            "original_price": 5090000,
            "price": 3110000,
            "selected": true,
            "seller": {
                "id": 1,
                "name": "Tiki Trading"
            },
            "sku": "9708011453693",
            "thumbnail_url": "https://salt.tikicdn.com/cache/280x280/ts/product/02/59/95/3897103ae79ab6e504d5178dbe2452b4.png"
        },
        {
            "child_id": 276614500,
            "discount_rate": 39,
            "id": 276615092,
            "images": [
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/32/1b/c1/44a9ba148ed3534a8f16ee8358f6084f.jpg",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/32/1b/c1/44a9ba148ed3534a8f16ee8358f6084f.jpg",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/32/1b/c1/44a9ba148ed3534a8f16ee8358f6084f.jpg"
                },
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/85/1b/e2/cc9e33d5f3c033b1bbd1a74469c28383.jpg",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/85/1b/e2/cc9e33d5f3c033b1bbd1a74469c28383.jpg",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/85/1b/e2/cc9e33d5f3c033b1bbd1a74469c28383.jpg"
                },
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/b2/36/d7/c933f95cc5f63f6f9458096bb4d295ca.jpg",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/b2/36/d7/c933f95cc5f63f6f9458096bb4d295ca.jpg",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/b2/36/d7/c933f95cc5f63f6f9458096bb4d295ca.jpg"
                },
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/b6/4d/fc/235b018d317ff1a4fbb21715aee63cd9.jpg",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/b6/4d/fc/235b018d317ff1a4fbb21715aee63cd9.jpg",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/b6/4d/fc/235b018d317ff1a4fbb21715aee63cd9.jpg"
                },
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/cd/e0/ec/1c8ad50e0ca4d4dbe08d6a1baaf7e2b6.jpg",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/cd/e0/ec/1c8ad50e0ca4d4dbe08d6a1baaf7e2b6.jpg",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/cd/e0/ec/1c8ad50e0ca4d4dbe08d6a1baaf7e2b6.jpg"
                }
            ],
            "inventory_status": "available",
            "name": "Điện Thoại Samsung Galaxy A16 LTE (4GB/128GB) - Đã kích hoạt bảo hành điện tử - Hàng Chính Hãng - Đen",
            "option1": "Đen",
            "original_price": 5090000,
            "price": 3110000,
            "selected": false,
            "seller": {
                "id": 1,
                "name": "Tiki Trading"
            },
            "sku": "4540652451483",
            "thumbnail_url": "https://salt.tikicdn.com/cache/280x280/ts/product/72/1f/59/c135735178fe2d653f2f14643fa4f18f.jpg"
        },
        {
            "child_id": 276617069,
            "discount_rate": 39,
            "id": 276617070,
            "images": [
                {
                    "large_url": "https://salt.tikicdn.com/cache/w1200/ts/product/8e/05/a3/13208d26745456f7a990dcbfa641ecf6.png",
                    "medium_url": "https://salt.tikicdn.com/cache/550x550/ts/product/8e/05/a3/13208d26745456f7a990dcbfa641ecf6.png",
                    "small_url": "https://salt.tikicdn.com/cache/200x280/ts/product/8e/05/a3/13208d26745456f7a990dcbfa641ecf6.png"
                }
            ],
            "inventory_status": "available",
            "name": "Điện Thoại Samsung Galaxy A16 LTE (4GB/128GB) - Đã kích hoạt bảo hành điện tử - Hàng Chính Hãng - Xám Bạc",
            "option1": "Xám Bạc",
            "original_price": 5090000,
            "price": 3110000,
            "selected": false,
            "seller": {
                "id": 1,
                "name": "Tiki Trading"
            },
            "sku": "3860522682789",
            "thumbnail_url": "https://salt.tikicdn.com/cache/280x280/ts/product/49/72/0e/4314d48b50bcff5a46bf7e5121f94d90.jpg"
        }
    ],
    "services_and_promotions": [],
    "promitions": [],
    "stock_item": {
        "max_sale_qty": 1000,
        "min_sale_qty": 1,
        "preorder_date": null,
        "qty": 1000
    },
    "quantity_sold": {
        "text": "Đã bán 497",
        "value": 497
    },
    "categories": {
        "id": 1795,
        "name": "Điện thoại Smartphone",
        "is_leaf": true
    },
    "breadcrumbs": [
        {
            "url": "/dien-thoai-may-tinh-bang/c1789",
            "name": "Điện Thoại - Máy Tính Bảng",
            "category_id": 1789
        },
        {
            "url": "/dien-thoai-smartphone/c1795",
            "name": "Điện thoại Smartphone",
            "category_id": 1795
        },
        {
            "url": "https://tiki.vn/dien-thoai-samsung-galaxy-a16-8gb-128gb-da-kich-hoat-bao-hanh-dien-tu-hang-chinh-hang-p276614493.html",
            "name": "Điện Thoại Samsung Galaxy A16 LTE (4GB/128GB) - Đã kích hoạt bảo hành điện tử - Hàng Chính Hãng",
            "category_id": 0
        }
    ],
    "installment_info_v2": null,
    "installment_info_v3": [
        {
            "display_text": "Ưu đãi đến 600k với thẻ TikiCard",
            "icon": "https://salt.tikicdn.com/ts/upload/73/4d/f7/f86e767bffc14aa3d6abed348630100b.png",
            "summary": "Ưu đãi đến 600k với thẻ TikiCard.",
            "title": "Ưu đãi đến 600k với thẻ TikiCard",
            "type": "other",
            "url": "https://tiki.vn/san-pham-so/the-dong-thuong-hieu-tiki-co-card"
        },
        {
            "display_text": "Mua trả góp từ 259.166<sup><small>₫</small></sup>/tháng",
            "icon": "https://salt.tikicdn.com/ts/upload/39/c5/e5/da087b9fd6fa5cf4c5aa0d6d2eb454df.png",
            "summary": "Qua thẻ tín dụng Visa, Mastercard, JCB.",
            "title": "Trả góp từ 259.166 <sup><small>₫</small></sup>/tháng",
            "type": "installment",
            "url": "https://tiki.vn/installment?masterId=276614493&&spid=276615090"
        },
        {
            "display_text": "Mua trước trả sau",
            "icon": "https://salt.tikicdn.com/ts/upload/2a/27/6a/7bbba1f6c93a1a42a3c314e7b5825f4c.png",
            "summary": "Duyệt hồ sơ nhanh chóng qua công ty tài chính.",
            "title": "Mua trước trả sau",
            "type": "postpaid",
            "url": "https://tiki.vn/mua-truoc-tra-sau/thanh-toan?src=pdp&masterId=276614493&pid=276615090&seller_id=1&qty=1"
        }
    ],
    "is_seller_in_chat_whitelist": true,
    "inventory": {
        "product_virtual_type": null,
        "fulfillment_type": "tiki_delivery"
    },
    "warranty_info": [
        {
            "name": "Thời gian bảo hành",
            "value": "12 Tháng"
        },
        {
            "name": "Hình thức bảo hành",
            "value": "Điện tử"
        },
        {
            "name": "Nơi bảo hành",
            "value": "Bảo hành chính hãng"
        },
        {
            "name": "Hướng dẫn bảo hành",
            "value": "Xem chi tiết",
            "url": "https://hotro.tiki.vn/s/article/chinh-sach-bao-hanh-tai-tiki-nhu-the-nao"
        }
    ],
    "return_and_exchange_policy": "Đổi trả trong<br><b>7 ngày</b><br>nếu sp lỗi.",
    "is_tier_pricing_available": false,
    "is_tier_pricing_eligible": false,
    "benefits": [
        {
            "icon": "https://salt.tikicdn.com/ts/upload/c5/37/ee/76c708d43e377343e82baee8a0340297.png",
            "text": "Được đồng kiểm khi nhận hàng"
        },
        {
            "icon": "https://salt.tikicdn.com/ts/upload/ea/02/b4/b024e431ec433e6c85d4734aaf35bd65.png",
            "text": "<b>Được hoàn tiền 200%</b> nếu là hàng giả."
        },
        {
            "icon": "https://salt.tikicdn.com/ts/upload/d8/c7/a5/1cd5bd2f27f9bd74b2c340b8e27c4d82.png",
            "text": "Đổi trả miễn phí trong 30 ngày. Được đổi ý.",
            "sub_text": [],
            "cta": {
                "text": "Chi tiết",
                "type": "open_policy_return_modal"
            }
        }
    ],
    "return_policy": {
        "body": [
            {
                "label": "",
                "content": [
                    "Được đổi ý (sản phẩm phải còn nguyên hộp, tem, phụ kiện, chưa kích hoạt bảo hành, không áp dụng đơn hàng trả góp), hoặc",
                    "Sản phẩm không đúng cam kết (lỗi kỹ thuật, giao sai/thiếu, bể vỡ…)"
                ]
            }
        ],
        "cta": {
            "link": "https://hotro.tiki.vn/s/article/chinh-sach-doi-tra-san-pham",
            "text_button": "Tìm hiểu thêm"
        },
        "title": "Đổi trả miễn phí trong 30 ngày"
    }
}
`
