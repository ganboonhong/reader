"use strict"

class Article {
    constructor() {
        this.daterangepickerId = "daterangepicker";

        this._initDateRangePicker(); // _initMainTable's dependency
        this._initArticleSourceSelect2();
        this._initMainTable();
        this._bindMenuButton();
        this._bindSubmitSourceFilter();

        this.table;
    }

    _initMainTable() {
        const _t = this;
        let draw = 1;
        this.table = $('#mainTable').DataTable({
            processing: true,
            serverSide: true,
            ordering: false,
            searching: false,
            lengthChange: false,
            ajax: {
                url: "get_article",
                data: function(dt) {
                    return {
                        draw: ++draw,
                        s_date: $(`#${_t.daterangepickerId}`).data('daterangepicker').startDate.format('YYYY-MM-DD'),
                        e_date: $(`#${_t.daterangepickerId}`).data('daterangepicker').endDate.format('YYYY-MM-DD'),
                        article_sources: _t._getSourceFilter()
                    };
                }
            },
            columnDefs: [

                {
                    targets: ['title_th'],
                    render: function(data, type, row) {
                        return _t._getTitleEle(data, row)
                    },
                }, {
                    targets: ['descr_th'],
                    render: function(data, type, row) {
                        return data
                    },
                }, {
                    targets: ['published_th'],
                    render: function(data, type, row) {
                        return _t._getPublishedAtEle(data)
                    },
                }
            ],
            columns: [

                {
                    data: "title",
                }, {
                    data: "description",
                }, {
                    data: "published_at",
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
        const $select2 = $('#sidebar-wrapper .article_source');
        $select2.select2({
            placeholder: "Please select a source",
            width: "95%",
            multiple: true,
        });
        $select2.val('').change();
        $('#sidebar-wrapper #comprehensive').val('cnn').change(); // default source
    }

    _bindMenuButton() {
        $("#menu-toggle").click(function(e) {
            e.preventDefault();
            $("#wrapper").toggleClass("toggled");
        });
    }

    _bindSubmitSourceFilter() {
        const _t = this;
        $("#sidebar-wrapper").on("click", "#submit_source", function() {
            if (!_t._getSourceFilter().length) {
                toastr.warning("Please select a source")
                return;
            }
            _t.table.ajax.reload();
        })
    }

    _getPublishedAtEle(data) {
        return moment(data).add(8, 'hours').format('YYYY-MM-DD HH:mm:ss');
    }

    _getSourceFilter() {
        let sources = [];
        $("#sidebar-wrapper .article_source").each(function() {
            const selected = $(this).val();
            sources.push(...selected)
        });

        return sources;
    }

    _getTitleEle(data, row) {
        let str = `[<b>${row.Source.name}</b>] `;
        str += data;
        return str;
    }
}