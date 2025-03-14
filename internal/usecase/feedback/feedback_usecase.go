package feedback

import (
	"context"

	"github.com/probuborka/feedback/internal/entity"
)

type feedback struct {
}

func NewFeedbackUseCase() feedback {
	return feedback{}
}

func (s feedback) AddFeedback(ctx context.Context, feedback entity.Feedback) error {

	//validate
	err := feedback.Validate()
	if err != nil {
		return err
	}

	// //search for recommendations in cache
	// recommendationCache, err := s.cache.FindByID(ctx, userRecommendationRequest.UserID)
	// if err != nil {
	// 	return "", err
	// }

	// //check recommendation from cache and user recommendation from request
	// if userRecommendationRequest.UserID == recommendationCache.UserID &&
	// 	userRecommendationRequest.UserName == recommendationCache.UserName &&
	// 	userRecommendationRequest.UserData.Profile == recommendationCache.UserData.Profile &&
	// 	userRecommendationRequest.UserData.Goals == recommendationCache.UserData.Goals &&
	// 	recommendationCache.Recommendations != "" {
	// 	return recommendationCache.Recommendations, nil
	// }

	// //get recommendations from AI
	// recommendations, err := s.ai.Recommendation(userRecommendationRequest)
	// if recommendations == "" || err != nil {
	// 	return "", err
	// }

	// userRecommendationRequest.Recommendations = recommendations

	// //save recommendations in cache
	// err = s.cache.Save(ctx, userRecommendationRequest)
	// if err != nil {
	// 	return "", err
	// }

	return nil
}
