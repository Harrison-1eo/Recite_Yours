<!-- src/views/AddWord.vue -->
<template>
    <div class="container mt-5">
        <button class="btn btn-primary mb-3" @click="$router.push('/user')">返回</button>
        <br>
        <h2>添加生词</h2>
        <form @submit.prevent="submitWord">
            <div class="mb-3">
                <label for="word" class="form-label">生词 (必填)</label>
                <input type="text"
                       v-model="word"
                       class="form-control"
                       id="word"
                       @focus="showSuggestions = true"
                       @blur="hideSuggestions"
                       required />
            </div>
            <div class="autocomplete-list"
                 v-if="suggestions !== undefined && suggestions.length > 0 && showSuggestions ">
                <ul class="list-group">
                    <li
                            class="list-group-item"
                            v-for="suggestion in suggestions"
                            :key="suggestion"
                            @click="selectSuggestion(suggestion)"
                    >
                        {{ suggestion }}
                    </li>
                </ul>
            </div>
            <div class="mb-3">
                <label for="translation" class="form-label">解释 (选填)</label>
                <input type="text"
                       v-model="translation"
                       class="form-control"
                       id="translation"
                       @click="fetchMeanings"
                       @focus="showMeanings = true"
                       @blur="hideMeanings"/>
            </div>
            <div class="meaning-list"
                 v-if="meanings !== undefined && meanings.length > 0 && showMeanings ">
                <ul class="list-group">
                    <li
                            class="list-group-item"
                            v-for="m in meanings"
                            :key="m"
                            @click="selectMeaning(m)"
                    >
                        {{ m.meaning }}
                    </li>
                </ul>
            </div>
            <div class="mb-3">
                <label for="exampleSentence" class="form-label">例句 (选填)</label>
                <input type="text" v-model="exampleSentence" class="form-control" id="exampleSentence" />
            </div>
            <div class="mb-3">
                <label for="note" class="form-label">个人备注 (选填)</label>
                <input type="text" v-model="note" class="form-control" id="note" />
            </div>
            <div class="mb-3">
                <label for="source" class="form-label">来源</label>
                <select v-model="selectedSource" class="form-select" id="source">
                    <option value=0>未选择来源</option>
                    <option v-for="source in sources" :key="source.ID" :value="source.ID">{{ source.name }}</option>
                </select>
                <button type="button" class="btn btn-link" data-bs-toggle="modal" data-bs-target="#addSourceModal">新建来源</button>
            </div>
            <button type="submit" class="btn btn-success">提交</button>
        </form>

        <!-- 新建来源的模态框 -->
        <div class="modal fade" id="addSourceModal" tabindex="-1" aria-labelledby="addSourceModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h1 class="modal-title fs-5" id="addSourceModalLabel">新建来源</h1>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                        <input type="text" v-model="newSource" class="form-control" placeholder="输入新来源" />
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
                        <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="addSource">添加来源</button>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
import axios_instance, {API} from '../utils/axios_config';
import Modal from "bootstrap/js/src/modal";
import {showSuccessToast} from "@/utils/showMessage";


export default {
    name: 'AddWord',
    data() {
        return {
            word: '',
            translation: '',
            exampleSentence: '',
            note: '',
            sources: [],
            selectedSource: 0,
            newSource: '',
            suggestions: [], // 用于存储联想提示词
            showSuggestions: false, // 控制提示框显示
            meanings: [], // 用于存储联想提示词
            showMeanings: false, // 控制提示框显示
        };
    },
    watch: {
        word(newValue) {
            if (newValue) {
                this.fetchSuggestions(newValue);
            } else {
                this.suggestions = []; // 清空提示
            }
        },
    },
    mounted() {
        this.fetchSources();
    },
    methods: {
        hideSuggestions() {
            // 延迟清除，以便用户在点击选项时能顺利选择
            setTimeout(() => {
                this.showSuggestions = false;
            }, 100);
        },
        hideMeanings() {
            // 延迟清除，以便用户在点击选项时能顺利选择
            setTimeout(() => {
                this.showMeanings = false;
            }, 100);
        },
        selectSuggestion(suggestion) {
            this.word = suggestion; // 选中提示词
            this.suggestions = [];  // 清空提示
        },
        selectMeaning(meaning) {
            this.translation = meaning.meaning; // 选中提示词
            this.meanings = [];  // 清空提示
        },
        async fetchSuggestions(prefix) {
            if (!prefix) return;
            try {
                this.suggestions = await axios_instance.get(API.WORDS_AUTOCOMPLETE + '?prefix=' + prefix);
            } catch (error) {
                console.error('Failed to fetch suggestions:', error);
            }
        },
        async fetchMeanings() {
            if (!this.word) return;
            try {
                this.meanings = await axios_instance.get(API.WORDS_MEANINGS + '?word=' + this.word);
            } catch (error) {
                console.error('Failed to fetch meanings:', error);
            }
        },
        async fetchSources() {
            try {
                this.sources = await axios_instance.get(API.SOURCES_LIST); // 假设返回的是一个数组
            } catch (error) {
                console.error('Failed to fetch sources:', error);
            }
        },
        async addSource() {
            if (this.newSource) {
                try {
                    await axios_instance.post(API.SOURCES_ADD, { name: this.newSource });
                    this.newSource = ''; // 清空输入框
                    await this.fetchSources(); // 重新获取来源列表
                    showSuccessToast('新来源已成功添加');
                } catch (error) {
                    console.error('Failed to add source:', error);
                }
            }
            // const modal = new Modal(document.getElementById('addSourceModal')); // 获取模态框实例
            // modal.hide(); // 关闭模态框
        },
        async submitWord() {
            try {
                const wordData = {
                    word: this.word,
                    translation: this.translation,
                    exampleSentence: this.exampleSentence,
                    note: this.note,
                    sourceID: this.selectedSource,
                };
                await axios_instance.post(API.WORDS_ADD, wordData);
                showSuccessToast('生词已成功添加');
                this.resetForm(); // 重置表单
            } catch (error) {
                console.error('Failed to submit word:', error);
            }
        },
        resetForm() {
            this.word = '';
            this.translation = '';
            this.exampleSentence = '';
            this.note = '';
        }
    }
}
</script>

<style scoped>
/* 限制模态框的最大宽度 */
.modal-dialog{
    max-width: 600px; /* 你可以根据需要调整这个值 */
}

.form-control{
    max-width: 600px; /* 你可以根据需要调整这个值 */
}

.form-select{
    max-width: 600px; /* 你可以根据需要调整这个值 */
}

/* 限制输入框的最大宽度 */
.modal-body {
    max-width: 100%;
}

.autocomplete-list {
    position: absolute;
    z-index: 1000;
    background-color: white;
    border: 1px solid #ccc;
    max-height: 200px; /* 调整高度 */
    overflow-y: auto;  /* 超出时显示滚动条 */
}

.meaning-list {
    max-width: 300px; /* 调整宽度 */
    position: absolute;
    z-index: 1000;
    background-color: white;
    border: 1px solid #ccc;
    max-height: 200px; /* 调整高度 */
    overflow-y: auto;  /* 超出时显示滚动条 */
}

.list-group-item {
    cursor: pointer;
}

.list-group-item:hover {
    background-color: #f0f0f0; /* 鼠标悬停效果 */
}


</style>