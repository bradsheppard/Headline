import { TopicServiceClient } from '../proto/topic/TopicServiceClientPb';
import {
    AddTopicsRequest, GetTopicsRequest, RemoveTopicsRequest, Topic,
} from '../proto/topic/topic_pb';
import { API_HOST } from './constants';

// eslint-disable-next-line
class TopicService {
    static async getTopics(userId: number): Promise<Topic[]> {
        const topicServiceClient = this.getClient();

        const request = new GetTopicsRequest();
        request.setUserid(userId);

        const response = await topicServiceClient.getTopics(request, null);
        return response.getTopicsList();
    }

    static async createTopic(topic: string): Promise<void> {
        const topicServiceClient = this.getClient();

        const request = new AddTopicsRequest();

        const createTopic = new Topic();
        createTopic.setName(topic);

        request.setTopicsList([createTopic]);
        request.setUserid(1);

        await topicServiceClient.addTopics(request, null);
    }

    static async deleteTopic(topic: string, userId: number): Promise<void> {
        const topicServiceClient = this.getClient();

        const request = new RemoveTopicsRequest();
        request.setTopicnamesList([topic]);
        request.setUserid(userId);

        await topicServiceClient.removeTopics(request, null);
    }

    private static getClient(): TopicServiceClient {
        return new TopicServiceClient(`http://${API_HOST}:80`);
    }
}

export default TopicService;
