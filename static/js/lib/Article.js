"use strict"

class Article {
    constructor() {

        $('#example').DataTable({
            responsive: true,
            columnDefs: [{
                responsivePriority: 1,
                targets: -1
            }]
        });

        // $('#mainTable').DataTable({
        //     responsive: true,
        //     columnDefs: [

        //         {
        //             responsivePriority: 1,
        //             targets: ['title_th'],
        //         }, {
        //             responsivePriority: 4,
        //             targets: ['descr_th'],
        //         }, {
        //             responsivePriority: 3,
        //             targets: ['published_th'],
        //         }, {
        //             responsivePriority: 2,
        //             targets: ['action_th'],
        //         }
        //     ]
        // });
        this.daterangepickerId = "daterangepicker";
        this.table;
        this.newsType = Article.CONSTANT.NEWS_TYPE.EVERYTHING;
        // this.newsType = Article.CONSTANT.NEWS_TYPE.TOPHEADLINE;

        this._initDateRangePicker(); // _initMainTable's dependency
        this._initArticleSourceSelect2();
        this._initTopHeadlineSelect2();
        this._initMainTable();
        this._bindMenuButton();
        this._bindSubmitSourceFilter();
        this._bindSubmitTopHeadlineFilter();
    }

    static get CONSTANT() {
        return {
            NEWS_TYPE: {
                EVERYTHING: "everything",
                TOPHEADLINE: "topheadline",
            }
        }
    }

    _initMainTable() {
        const _t = this;
        let draw = 1;
        this.table = $('#mainTable').DataTable({
            processing: true,
            serverSide: true,
            responsive: true,
            ordering: false,
            searching: false,
            lengthChange: false,
            ajax: {
                url: "get_article",
                data: function (dt) {
                    return {
                        draw: ++draw,
                        s_date: $(`#${_t.daterangepickerId}`).data('daterangepicker').startDate.format('YYYY-MM-DD'),
                        e_date: $(`#${_t.daterangepickerId}`).data('daterangepicker').endDate.format('YYYY-MM-DD'),
                        article_sources: _t._getSourceFilter(),
                        country: _t._getTopHeadlineFilter(),
                        page: $('#mainTable').DataTable().page(),
                        dt: dt,
                        news_type: _t.newsType
                    };
                }
            },
            columnDefs: [

                {
                    targets: ['title_th'],
                    responsivePriority: 1,
                    width: "40%",
                    render: function (data, type, row) {
                        return _t._getTitleEle(data, row)
                    },
                }, {
                    targets: ['descr_th'],
                    responsivePriority: 2,
                    width: "60%",
                    render: function (data, type, row) {
                        return data
                    },
                }
            ],
            // columnDefs: [

            //     {
            //         responsivePriority: 1,
            //         targets: ['title_th'],
            //     }, {
            //         responsivePriority: 4,
            //         targets: ['descr_th'],
            //     }, {
            //         responsivePriority: 3,
            //         targets: ['published_th'],
            //     }, {
            //         responsivePriority: 2,
            //         targets: ['action_th'],
            //     }
            // ],
            columns: [

                {
                    data: "title",
                }, {
                    data: "description",
                }
            ]
        });
    }

    _initDateRangePicker() {
        const _t = this;
        const $daterangepicker = $(`#${this.daterangepickerId}`)
        $daterangepicker.daterangepicker({
            startDate: moment().subtract(1, 'day'),
            locale: {
                format: "DD/MM/YY"
            }
        });
    }

    _initArticleSourceSelect2() {
        const _t = this;
        const $select2 = $('#sidebar-wrapper .article_source');
        $select2.select2({
            placeholder: "Select source",
            width: "100%",
            multiple: true,
        });
        $select2.val('').change();
        // $('#sidebar-wrapper #countries').val('tw').change(); // default source
        // $('#sidebar-wrapper #countries').val('us').change(); // default source
        $('#sidebar-wrapper #tech').val('techcrunch').change(); // default source
        $('#sidebar-wrapper #comprehensive').val('cnn').change(); // default source
        $('#sidebar-wrapper #business').val('business-insider').change(); // default source
        $('#sidebar-wrapper #sports').val('bleacher-report').change(); // default source
    }

    _initTopHeadlineSelect2() {
        const _t = this;
        // $('#sidebar-wrapper #tech').val('techcrunch').change(); // default source
    }

    _initTopHeadlineSelect2() {
        const _t = this;
        const $select2 = $('#sidebar-wrapper .top_headline_select2');
        $select2.select2({
            placeholder: "Select country",
            width: "85%",
        });
    }

    _hideFilterModal() {
        $("#filterModal").modal('hide');
    }

    _bindMenuButton() {
        $("#menu-toggle").click(function (e) {
            e.preventDefault();
            $("#wrapper").toggleClass("toggled");
        });
    }

    _bindSubmitTopHeadlineFilter() {
        const _t = this;
        $("#sidebar-wrapper").on("click", "#submit_top_headline_filter", function () {
            if (!_t._getTopHeadlineFilter()) {
                toastr.warning("Please select country")
                return;
            }
            _t.newsType = Article.CONSTANT.NEWS_TYPE.TOPHEADLINE;
            _t.table.ajax.reload();
            _t._hideFilterModal();
        })
    }

    _bindSubmitSourceFilter() {
        const _t = this;
        $("#sidebar-wrapper").on("click", "#submit_source", function () {
            if (!_t._getSourceFilter().length) {
                toastr.warning("Please select a source")
                return;
            }
            _t.newsType = Article.CONSTANT.NEWS_TYPE.EVERYTHING
            _t.table.ajax.reload();
            _t._hideFilterModal();
        })
    }

    _getSourceFilter() {
        let sources = [];
        $("#sidebar-wrapper .article_source").each(function () {
            const selected = $(this).val();
            sources.push(...selected)
        });

        return sources;
    }

    _getTopHeadlineFilter() {
        return $("#sidebar-wrapper .top_headline_select2").val();
    }

    _getTitleEle(data, row) {
        const publishedAt = moment(row.publishedAt).fromNow();
		let imageTd = '';
		if (row.urlToImage) {
			imageTd = `
			<td width="140px" class="title-td">
				<a href="${row.url}" target="_blank" title="Open in new tab">
					<img style="display: block; " width="140px" src="${row.urlToImage}" />
				</a>
			</td>`;
		}
        return `
		<table>
			<tr class="title-tr">
				${imageTd}
				<td class="title-td">
					<a href="${row.url}" target="_blank" title="Open in new tab">
						<i class="fas fa-external-link-alt fa"></i>
						<span style="color: black;">${data}</span>
					</a>
					<span style="color: #8c8c8c; font-size: 12px;">
						<br>
						${row.Source.name}  |  ${publishedAt}
					</span>
				</td>
			</tr>
		</table>
		`;
    }
}
