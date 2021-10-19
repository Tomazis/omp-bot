package review

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/service"
)

type ReviewService interface {
	Describe(reviewID uint64) (*service.Review, error)
	List(cursor uint64, limit uint64) ([]service.Review, error)
	Create(review service.Review) (uint64, error)
	Update(reviewID uint64, review service.Review) error
	Remove(reviewID uint64) (bool, error)
}

type DummyReviewService struct {
}

func NewDummyReviewService() *DummyReviewService {
	return &DummyReviewService{}
}

func (s *DummyReviewService) Describe(reviewID uint64) (*service.Review, error) {
	reviewID--
	if reviewID < uint64(len(service.Reviews)) {
		return &service.Reviews[reviewID], nil
	}
	return nil, errors.New("Index out of range")
}

func (s *DummyReviewService) List(cursor uint64, limit uint64) ([]service.Review, error) {
	if cursor < uint64(len(service.Reviews)) {
		if cursor+limit > uint64(len(service.Reviews)) {
			return service.Reviews[cursor:], nil
		}
		return service.Reviews[cursor : cursor+limit], nil
	}
	return nil, errors.New("Index out of range")
}

func (s *DummyReviewService) Create(review service.Review) (uint64, error) {
	service.Reviews = append(service.Reviews, review)
	return uint64(len(service.Reviews)), nil
}

func (s *DummyReviewService) Update(reviewID uint64, review service.Review) error {
	return errors.New("Not implemented")
}

func (s *DummyReviewService) Remove(reviewID uint64) (bool, error) {
	reviewID--
	if reviewID < uint64(len(service.Reviews)) {
		service.Reviews = append(service.Reviews[:reviewID], service.Reviews[reviewID+1:]...)
		return true, nil
	}
	return false, errors.New("Index out of range")
}
