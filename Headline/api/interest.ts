import { InterestServiceClient } from '../proto/interest/InterestServiceClientPb';
import {
    AddInterestsRequest,
    CreateInterest,
    DeleteInterestsRequest,
    GetInterestsRequest,
    type Interest,
} from '../proto/interest/interest_pb';
import { API_HOST } from './constants';

// eslint-disable-next-line
class InterestService {
    static async getInterests(userId: number): Promise<Interest[]> {
        const interestServiceClient = this.getClient();

        const request = new GetInterestsRequest();
        request.setUserid(userId);

        const response = await interestServiceClient.getInterests(request, null);
        return response.getInterestsList();
    }

    static async createInterest(interest: string): Promise<void> {
        const interestServiceClient = this.getClient();

        const request = new AddInterestsRequest();
        const createInterest = new CreateInterest();

        createInterest.setName(interest);
        createInterest.setUserid(1);

        request.setInterestsList([createInterest]);
        request.setUserid(1);

        await interestServiceClient.addInterests(request, null);
    }

    static async deleteInterest(id: number): Promise<void> {
        const interestServiceClient = this.getClient();

        const request = new DeleteInterestsRequest();
        request.setIdsList([id]);

        await interestServiceClient.deleteInterests(request, null);
    }

    private static getClient(): InterestServiceClient {
        return new InterestServiceClient(`http://${API_HOST}:80`);
    }
}

export default InterestService;
