<template>
    <div class="container mt-4">
        <!-- 返回按钮 -->
        <button class="btn btn-primary mb-3" @click="goToHome">返回</button>

        <!-- 选项框 -->
        <div class="mb-3">
            <label for="startDate" class="form-label">起始日期：</label>
            <input type="date" id="startDate" class="form-control" v-model="startDate">
        </div>
        <div class="mb-3">
            <label for="endDate" class="form-label">终止日期：</label>
            <input type="date" id="endDate" class="form-control" v-model="endDate">
        </div>
        <div class="mb-3">
            <button class="btn btn-outline-primary me-2" @click="selectToday">选择今天</button>
            <button class="btn btn-outline-primary me-2" @click="selectYesterday">选择昨天</button>
            <button class="btn btn-outline-primary" @click="selectDayBeforeYesterday">选择前天</button>
        </div>
        <div class="mb-3">
            <label for="source" class="form-label">来源列表：</label>
            <select id="source" class="form-select" v-model="selectedSource">
                <option value=0>全部来源</option>
                <option v-for="source in sources" :key="source.ID" :value="source.ID">{{ source.name }}</option>
            </select>
        </div>
        <div class="mb-3 form-check">
            <input type="checkbox" id="showDefinitions" class="form-check-input" v-model="showDefinitions">
            <label for="showDefinitions" class="form-check-label">显示单词解释</label>
        </div>

        <div class="mt-4">
            <button class="btn btn-primary me-3" @click="fetchWords">获取单词</button>
            <button class="btn btn-secondary me-3" @click="shuffleWords">打乱顺序</button>
            <button class="btn btn-success" @click="exportWords">导出单词</button>
        </div>

        <!-- 展示单词列表 -->
        <div v-if="words.length > 0">
            <h4 class="mt-4 mb-3">单词列表</h4>
            <table class="table table-hover">
                <thead>
                    <tr>
                        <th scope="col">序号</th>
                        <th scope="col">单词</th>
                        <th scope="col">解释</th>
                        <th scope="col">注释</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(word, index) in words" :key="word.ID">
                        <th scope="row" style="width: 70px">{{ index + 1 }}</th>
                        <td>{{ word.word }}</td>
                        <td><div v-if="showDefinitions">{{ word.translation }}</div></td>
                        <td><div v-if="showDefinitions">{{ word.note }}</div></td>
                    </tr>
                </tbody>
            </table>
        </div>


        <!-- 按钮 -->

    </div>
</template>

<script>
import axios_instance, {API} from "@/utils/axios_config";

export default {
    data() {
        return {
            startDate: '',
            endDate: '',
            selectedSource: 0,
            showDefinitions: false,
            words: [], // 单词列表，从后端接口获取
            sources: [],
        };
    },
    created() {
        this.selectToday();
        this.fetchSources();
    },
    methods: {
        async fetchSources() {
            try {
                this.sources = await axios_instance.get(API.SOURCES_LIST); // 假设返回的是一个数组
            } catch (error) {
                console.error('Failed to fetch sources:', error);
            }
        },
        goToHome() {
            this.$router.push({ name: 'UserHome' });
        },
        async fetchWords() {
            try {
                this.words = await axios_instance.get(API.WORDS_LIST, {
                    params: {
                        start_time: this.startDate,
                        end_time: this.endDate,
                        source_id: this.selectedSource,
                    },
                });
            } catch (error) {
                console.error('Failed to fetch words:', error);
            }
        },
        shuffleWords() {
            this.words = this.shuffleArray(this.words);
        },
        shuffleArray(array) {
            // 使用 Fisher-Yates 算法进行数组元素随机排序
            let currentIndex = array.length, temporaryValue, randomIndex;

            // While there remain elements to shuffle...
            while (0 !== currentIndex) {

                // Pick a remaining element...
                randomIndex = Math.floor(Math.random() * currentIndex);
                currentIndex -= 1;

                // And swap it with the current element.
                temporaryValue = array[currentIndex];
                array[currentIndex] = array[randomIndex];
                array[randomIndex] = temporaryValue;
            }

            return array;
        },
        exportWords() {
            // 导出当前单词的方法
            console.log('导出当前单词');
        },
        selectToday() {
            this.startDate = this.formatDate(new Date());
            this.endDate = this.startDate;
        },
        selectYesterday() {
            const yesterday = new Date();
            yesterday.setDate(yesterday.getDate() - 1);
            this.startDate = this.formatDate(yesterday);
            this.endDate = this.startDate;
        },
        selectDayBeforeYesterday() {
            const dayBeforeYesterday = new Date();
            dayBeforeYesterday.setDate(dayBeforeYesterday.getDate() - 2);
            this.startDate = this.formatDate(dayBeforeYesterday);
            this.endDate = this.startDate;
        },
        formatDate(date) {
            const year = date.getFullYear();
            let month = date.getMonth() + 1;
            let day = date.getDate();
            if (month < 10) {
                month = `0${month}`;
            }
            if (day < 10) {
                day = `0${day}`;
            }
            return `${year}-${month}-${day}`;
        },
    },
};
</script>

<style>
/* 可以添加自定义的样式 */
</style>
