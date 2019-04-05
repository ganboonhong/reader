"use strict"

class Article {
    constructor() {
        this.daterangepickerId = "daterangepicker";

        this._initDateRangePicker(); // _initMainTable's dependency
        this._initArticleSourceSelect2();
        this._initMainTable();
        this._bindMenuButton();

        this.table;
    }

    _initMainTable() {
        const _t = this;
        let draw = 1;
        this.table = $('#mainTable').DataTable({
            processing: true,
            serverSide: true,
            ordering: false,
            ajax: {
                url: "get_article",
                data: function(dt) {
                    return {
                        draw: ++draw,
                        s_date: $(`#${_t.daterangepickerId}`).data('daterangepicker').startDate.format('YYYY-MM-DD'),
                        e_date: $(`#${_t.daterangepickerId}`).data('daterangepicker').endDate.format('YYYY-MM-DD'),
                        article_sources: $('#sidebar-wrapper #article_source').val()
                    };
                }
            },
            columnDefs: [

                {
                    targets: ['title_th'],
                    render: function(data, type, row) {
                        return data
                    },
                }, {
                    targets: ['descr_th'],
                    render: function(data, type, row) {
                        return data
                    },
                }, {
                    targets: ['published_th'],
                    render: function(data, type, row) {
                        return data
                    },
                }, {
                    targets: ['content_th'],
                    render: function(data, type, row) {
                        return data
                    },
                },
            ],
            columns: [

                {
                    data: "title",
                }, {
                    data: "description",
                }, {
                    data: "publishedAt",
                }, {
                    data: "content",
                }
            ]
        });
    }

    _initDateRangePicker() {
        const _t = this;
        const $daterangepicker = $(`#${this.daterangepickerId}`)
        $daterangepicker.daterangepicker();

        $daterangepicker.on("apply.daterangepicker", function(e, picker) {
            _t.table.ajax.reload();
        })
    }

    _initArticleSourceSelect2() {
        const _t = this;
        const $select2 = $('#sidebar-wrapper #article_source');
        $select2.select2({
            width: "100%",
            multiple: true,
        });
        $select2.on("change", function(e) {
            _t.table.ajax.reload();
        })
    }

    _bindMenuButton() {
        $("#menu-toggle").click(function(e) {
            e.preventDefault();
            $("#wrapper").toggleClass("toggled");
        });
    }
}